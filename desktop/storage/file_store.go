package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"sync"
	"time"
)

type FileStore[F any] struct {
	lock      sync.Mutex
	path      string
	cache     F
	isChanged bool
	isSync    bool
	quit      chan struct{}
}

func NewFileStore[F any](path string, _ F) (*FileStore[F], error) {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			file, err2 := os.Create(path)
			if err2 != nil {
				return nil, err2
			}
			defer file.Close()

			fs := &FileStore[F]{
				path: path,
				quit: make(chan struct{}),
			}
			go fs.loopSync()
			return fs, nil
		}
		return nil, err
	}
	if stat.IsDir() {
		return nil, errors.New("path is a directory: " + path)
	}
	fs := &FileStore[F]{
		path: path,
		quit: make(chan struct{}),
	}
	err = fs.load()
	if err != nil {
		return nil, err
	}
	go fs.loopSync()
	return fs, nil
}

func (fs *FileStore[F]) load() error {
	fs.lock.Lock()
	defer fs.lock.Unlock()
	bytes, err := os.ReadFile(fs.path)
	if err != nil {
		return err
	}
	if len(bytes) == 0 {
		return nil
	}
	var cache F
	err = json.Unmarshal(bytes, &cache)
	if err != nil {
		return err
	}
	fs.cache = cache
	return nil
}

func (fs *FileStore[F]) loopSync() error {
	for {
		select {
		case <-fs.quit:
		case <-time.After(10 * time.Second):
			func() {
				fs.lock.Lock()
				defer fs.lock.Unlock()
				err := fs.sync()
				if err != nil {
					runtime.LogWarningf(context.Background(), "sync file %s failed: %v", fs.path, err)
				}
			}()
		}
	}
}

func (fs *FileStore[F]) sync() error {
	fs.isSync = true
	defer func() {
		fs.isSync = false
		fs.isChanged = false
	}()
	if !fs.isChanged {
		return nil
	}
	bytes, err := json.Marshal(fs.cache)
	if err != nil {
		return err
	}
	if len(bytes) == 4 {
		if string(bytes) == "null" {
			return nil
		}
	}
	err = os.WriteFile(fs.path, bytes, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStore[F]) Set(t F) {
	defer func() {
		fs.isChanged = true
	}()
	fs.cache = t

}

func (fs *FileStore[F]) Update(fn func(f F) F) {
	defer func() {
		fs.isChanged = true
	}()
	f := fn(fs.cache)
	fs.cache = f
}

func (fs *FileStore[F]) Get() F {
	return fs.cache
}

func (fs *FileStore[F]) IsSync() bool {
	return fs.isSync
}

func (fs *FileStore[F]) Close() {
	fs.lock.Lock()
	defer fs.lock.Unlock()
	if fs.isChanged {
		fs.sync()
	}
	close(fs.quit)
}

package service

import (
	"context"
	"path/filepath"
	"sync"
	"tapi/desktop/logger"
	"tapi/desktop/model"
	"tapi/desktop/storage"
)

type storeSvc struct {
	ctx                 context.Context
	db                  *storage.LocalStore
	projectMap          map[string]*model.Project
	environments        map[string]*model.Environment
	collections         map[string]*model.Collection
	collectionTree      map[string]*model.CollectionTree
	collectionTreeArray []*model.CollectionTree
	requests            map[string]*model.Request
}

var store *storeSvc
var onceStore sync.Once

func Store() *storeSvc {
	onceStore.Do(func() {
		store = &storeSvc{}
	})
	return store
}

func (c *storeSvc) Init(ctx context.Context) {
	logger.Info("===========================> local storage init")
	c.ctx = ctx
	rootDir := System().GetRootDir()
	dataDir := filepath.Join(rootDir, System().Preference.Get().Data.DataDir)
	db := storage.NewLocalStore(dataDir)
	err := db.Open()
	if err != nil {
		logger.Fatal(err)
	}
	c.db = db
}

func (c *storeSvc) Close() {
	if c.db != nil {
		err := c.db.Close()
		logger.Error(err)
	}
}

func (c *storeSvc) Get(ctx context.Context, key []byte, val storage.Storable, inNamespace ...bool) error {
	return c.db.Get(ctx, key, val, inNamespace...)
}

func (c *storeSvc) Set(ctx context.Context, key []byte, val storage.Storable, inNamespace ...bool) error {
	return c.db.Set(ctx, key, val, inNamespace...)
}

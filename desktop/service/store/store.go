package store

import (
	"context"
	"path/filepath"
	"sync"
	"tapi/desktop/logger"
	"tapi/desktop/model"
	"tapi/desktop/service"
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
	c.ctx = ctx
	rootDir := service.System().GetRootDir()
	dataDir := filepath.Join(rootDir, service.System().Preference.Get().Data.DataDir)
	db := storage.NewLocalStore(dataDir)
	c.db = db
}

func (c *storeSvc) Close() {
	if c.db != nil {
		err := c.db.Close()
		logger.Error(err)
	}
}

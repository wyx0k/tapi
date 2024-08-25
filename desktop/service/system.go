package service

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"tapi/desktop/logger"
	"tapi/desktop/model"
	"tapi/desktop/storage"
	"tapi/desktop/utils"
)

type systemSvc struct {
	rootDir    string
	Preference *storage.FileStore[*model.Preference]
	UIState    *storage.FileStore[*model.AppUIState]
}

var system *systemSvc
var onceSystem sync.Once

func System() *systemSvc {
	onceSystem.Do(func() {
		system = &systemSvc{}
	})
	return system
}

func (c *systemSvc) Init(ctx context.Context) {
	dir, err := os.Getwd()
	if err != nil {
		c.rootDir = "tapi"
	} else {
		c.rootDir = filepath.Join(dir, "tapi")
	}
	if err = utils.CheckOrCreateDir(c.rootDir); err != nil {
		logger.Fatal(ctx, "can not create directory at tapi")
	}
	if err = utils.CheckOrCreateDir(filepath.Join(c.rootDir, "data")); err != nil {
		logger.Fatal(ctx, "can not create data directory at tapi")
	}
	if err = utils.CheckOrCreateDir(filepath.Join(c.rootDir, "settings")); err != nil {
		logger.Fatal(ctx, "can not create setting directory at tapi")
	}
	c.Preference, err = storage.NewFileStore(filepath.Join(c.rootDir, "settings", "preference.toml"), &model.Preference{})
	if err != nil {
		logger.Fatal(ctx, "can not read settings file at tapi: %v", err)
	}
	c.UIState, err = storage.NewFileStore(filepath.Join(c.rootDir, "settings", "ui-state.toml"), &model.AppUIState{})
	if err != nil {
		logger.Fatal(ctx, "can not read settings file at tapi: %v", err)
	}
	if c.Preference.Get() == nil {
		c.Preference.Set(&model.Preference{Data: model.PreferencesData{DataDir: "data"}})
	}
	if c.UIState.Get() == nil {
		c.UIState.Set(&model.AppUIState{})
	}
}
func (c *systemSvc) Close() {
	c.Preference.Close()
	c.UIState.Close()
}

func (c *systemSvc) GetRootDir() string {
	return c.rootDir
}

func getDataDir() string {
	return System().Preference.Get().Data.DataDir
}

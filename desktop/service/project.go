package service

import (
	"context"
	"sync"
	"tapi/desktop/service/store"
)

type projectSvc struct{}

var project *projectSvc
var onceProject sync.Once

func Project() *projectSvc {
	onceProject.Do(func() {
		project = &projectSvc{}
	})
	return project
}

func (c *projectSvc) Init(ctx context.Context) {
	store.Store().db

}
func (c *projectSvc) Close() {

}

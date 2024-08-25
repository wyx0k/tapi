package service

import (
	"context"
	"sync"
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

}
func (c *projectSvc) Close() {

}

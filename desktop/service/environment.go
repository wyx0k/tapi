package service

import (
	"context"
	"sync"
)

type environmentSvc struct{}

var environment *environmentSvc
var onceEnvironment sync.Once

func Environment() *environmentSvc {
	onceEnvironment.Do(func() {
		environment = &environmentSvc{}
	})
	return environment
}

func (c *environmentSvc) Init(ctx context.Context) {

}
func (c *environmentSvc) Close() {

}

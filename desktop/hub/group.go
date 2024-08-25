package hub

import (
	"context"
	"sync"
)

type groupSvc struct{}

var group *groupSvc
var onceGroup sync.Once

func Group() *groupSvc {
	onceGroup.Do(func() {
		group = &groupSvc{}
	})
	return group
}

func (c *groupSvc) Init(ctx context.Context) {

}

func (c *groupSvc) Close() {

}

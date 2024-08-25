package hub

import (
	"context"
	"sync"
)

type hubSvc struct{}

var hub *hubSvc
var onceHub sync.Once

func Hub() *hubSvc {
	onceHub.Do(func() {
		hub = &hubSvc{}
	})
	return hub
}

func (c *hubSvc) Init(ctx context.Context) {

}

func (c *hubSvc) Close() {

}

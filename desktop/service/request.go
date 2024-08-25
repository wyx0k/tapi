package service

import (
	"context"
	"sync"
)

type requestSvc struct{}

var request *requestSvc
var onceRequest sync.Once

func Request() *requestSvc {
	onceRequest.Do(func() {
		request = &requestSvc{}
	})
	return request
}

func (c *requestSvc) Init(ctx context.Context) {

}
func (c *requestSvc) Close() {

}

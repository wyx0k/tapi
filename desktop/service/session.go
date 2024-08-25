package service

import (
	"context"
	"sync"
)

type sessionSvc struct {
}

var session *sessionSvc
var onceSession sync.Once

func Session() *sessionSvc {
	onceSession.Do(func() {
		session = &sessionSvc{}
	})
	return session
}

func (c *sessionSvc) Init(ctx context.Context) {

}

func (c *sessionSvc) Close() {

}

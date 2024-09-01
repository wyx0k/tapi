package service

import (
	"context"
	"sync"
	"tapi/desktop/logger"
	"tapi/desktop/model"
	"tapi/desktop/storage"
)

type sessionSvc struct {
	ctx     context.Context
	session *model.Session
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
	c.ctx = ctx
	err := Store().Get(ctx, storage.Session(), c.session, false)
	if err != nil {
		logger.Fatal("can not init session", err)
	}
	if c.session == nil {
		c.session = &model.Session{
			TabSession: model.TabSession{},
		}
	}
}

func (c *sessionSvc) Close() {
	if c.session != nil {
		err := Store().Set(c.ctx, storage.Session(), c.session, false)
		logger.Error("close session", err)
	}
}

func (c *sessionSvc) GetTabSession() model.TabSession {
	return c.session.TabSession
}

func (c *sessionSvc) SetTabSession(tabSession model.TabSession) {
	c.session.TabSession = tabSession
}

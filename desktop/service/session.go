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
		session = &sessionSvc{
			session: &model.Session{},
		}
	})
	return session
}

func (c *sessionSvc) Init(ctx context.Context) {
	c.ctx = ctx
	err := Store().Get(ctx, storage.Session(), c.session, false)
	if err != nil {
		logger.Fatal("can not init session", err)
	}
	logger.Info(c.session)
	if c.session.TabSession.TabScrollLeft == "" {
		c.session.TabSession.TabScrollLeft = "0"
	}
}

func (c *sessionSvc) Close() {
	if c.session != nil {
		err := Store().Set(c.ctx, storage.Session(), c.session, false)
		logger.Error("close session", err)
	}
}

func (c *sessionSvc) GetTabSession() model.Response {
	return model.Success(c.session.TabSession)
}

func (c *sessionSvc) SetTabSessionTabs(tabs []model.Tab) model.Response {
	c.session.TabSession.Tabs = tabs
	return model.Success("")
}
func (c *sessionSvc) SetTabSessionPosition(current model.Tab, scrollLeft string) model.Response {
	logger.Info(current, scrollLeft)
	c.session.TabSession.CurrentTab = current
	c.session.TabSession.TabScrollLeft = scrollLeft
	return model.Success("")
}

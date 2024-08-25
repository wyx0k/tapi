package hub

import "sync"

type userSvc struct{}

var user *userSvc
var onceUser sync.Once

func User() *userSvc {
	onceUser.Do(func() {
		user = &userSvc{}
	})
	return user
}

func (c *userSvc) Init() {

}

func (c *userSvc) Close() {

}

package service

import (
	"context"
	"sync"
)

type collectionSvc struct{}

var collection *collectionSvc
var onceCollection sync.Once

func Collection() *collectionSvc {
	onceCollection.Do(func() {
		collection = &collectionSvc{}
	})
	return collection
}

func (c *collectionSvc) Init(ctx context.Context) {

}
func (c *collectionSvc) Close() {

}

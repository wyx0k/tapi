package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/assert"
	"tapi/desktop/ns"
	"testing"
)

type TestStorable struct {
	Name string `json:"name"`
}

func (t *TestStorable) Marshal() []byte {
	b, _ := json.Marshal(t)
	return b
}

func (t *TestStorable) Unmarshal(bytes []byte) error {
	err := json.Unmarshal(bytes, t)
	return err
}

func TestLocalStore(t *testing.T) {
	s := NewLocalStore("../../test")
	s.Open()
	defer s.Close()
	ctx := ns.ChangeNamespace(context.Background(), "test")
	var p uint64 = 1
	var c uint64 = 1
	var e uint64 = 1
	var r uint64 = 1

	err := s.Set(ctx, Request(p, c, r), &TestStorable{Name: "r"})
	assert.Nil(t, err)
	err = s.Set(ctx, Collection(p, c), &TestStorable{Name: "c"})
	assert.Nil(t, err)
	err = s.Set(ctx, Project(p), &TestStorable{Name: "p"})
	assert.Nil(t, err)
	err = s.Set(ctx, Environment(p, e), &TestStorable{Name: "e"})
	assert.Nil(t, err)
	err = s.Set(ctx, Session(), &TestStorable{Name: "s"})
	assert.Nil(t, err)
	ts := &TestStorable{}
	err = s.Get(ctx, Project(p), ts)
	assert.Nil(t, err)
	assert.Equal(t, ts.Name, "p")

	err = s.Get(ctx, Environment(p, e), ts)
	assert.Nil(t, err)
	assert.Equal(t, ts.Name, "e")

	err = s.Get(ctx, Collection(p, c), ts)
	assert.Nil(t, err)
	assert.Equal(t, ts.Name, "c")

	err = s.Get(ctx, Request(p, c, r), ts)
	assert.Nil(t, err)
	assert.Equal(t, ts.Name, "r")

	s.db.View(func(txn *badger.Txn) error {
		_, err2 := txn.Get([]byte("v1.namespaces.test.projects.1.collections.1.requests.1"))
		assert.Nil(t, err2)
		return err2
	})

	s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			fmt.Println(string(item.Key()))
		}
		return nil
	})
}

func TestLocalStoreKey(t *testing.T) {
	keys := []StoreKey{
		StoreKey("v1.namespaces.test.projects.1.collections.1.requests.1"),
		StoreKey("v1.namespaces.test.projects.1.collections.1"),
		StoreKey("v1.namespaces.test.projects.1.environments.1"),
		StoreKey("v1.namespaces.test.projects.1"),
		StoreKey("v1.namespaces.test"),
	}
	for i, key := range keys {
		assert.Equal(t, false, key.IsSubKey())
		id, err := key.GetId()
		if i == 4 {
			assert.Equal(t, err, ErrWrongTypeForId)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, uint64(1), id)
		}
		ns, err := key.GetNamespace()
		assert.Nil(t, err)
		assert.Equal(t, ns, "test")
	}
	keys = []StoreKey{
		StoreKey(".projects.1.collections.1.requests.1"),
		StoreKey(".projects.1.collections.1"),
		StoreKey(".projects.1.environments.1"),
		StoreKey(".projects.1"),
	}
	for _, key := range keys {
		assert.Equal(t, true, key.IsSubKey())
		id, err := key.GetId()
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), id)
		ns, err := key.GetNamespace()
		assert.Equal(t, err, ErrKeyNotFound)
		assert.Equal(t, ns, "")
	}
}

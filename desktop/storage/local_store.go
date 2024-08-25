package storage

import (
	"bytes"
	"context"
	"github.com/dgraph-io/badger/v4"
	"tapi/desktop/logger"
	"tapi/desktop/ns"
)

type Storable interface {
	Marshal() []byte
	Unmarshal([]byte) error
}

type LocalStore struct {
	path string
	db   *badger.DB
}

func NewLocalStore(path string) *LocalStore {
	s := &LocalStore{path: path}
	return s
}

func (s *LocalStore) Open() error {
	opt := badger.DefaultOptions(s.path).
		WithLogger(logger.DefaultLogger).
		WithLoggingLevel(badger.WARNING)
	db, err := badger.Open(opt)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *LocalStore) Close() error {
	return s.db.Close()
}

func (s *LocalStore) Set(ctx context.Context, key []byte, val Storable) error {
	k, err := buildKey(ctx, key)
	if err != nil {
		return err
	}
	bytes := val.Marshal()
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(k, bytes)
	})
}

func (s *LocalStore) Get(ctx context.Context, key []byte, obj Storable) error {
	k, err := buildKey(ctx, key)
	if err != nil {
		return err
	}
	return s.db.View(func(txn *badger.Txn) error {

		item, err2 := txn.Get(k)
		if err2 != nil {
			return err2
		}
		return item.Value(func(val []byte) error {
			return obj.Unmarshal(val)
		})
	})
}

func buildKey(ctx context.Context, key []byte) ([]byte, error) {
	ns, err := ns.NamespaceMust(ctx)
	if err != nil {
		return []byte("."), err
	}
	keys := [][]byte{keyVersion, keyNamespaces, []byte(ns)}
	b := bytes.Join(keys, keySep)
	return append(b, key...), nil
}

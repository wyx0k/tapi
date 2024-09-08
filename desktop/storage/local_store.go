package storage

import (
	"bytes"
	"context"
	"errors"
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

func (s *LocalStore) Set(ctx context.Context, key []byte, val Storable, inNamespace ...bool) error {
	k, err := buildKey(ctx, key, inNamespace...)
	if err != nil {
		return err
	}
	bytes := val.Marshal()
	if bytes == nil {
		return nil
	}
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(k, bytes)
	})
}

func (s *LocalStore) Get(ctx context.Context, key []byte, obj Storable, inNamespace ...bool) error {
	k, err := buildKey(ctx, key, inNamespace...)
	if err != nil {
		return err
	}
	return s.db.View(func(txn *badger.Txn) error {

		item, err2 := txn.Get(k)
		if err2 != nil {
			if errors.Is(err2, badger.ErrKeyNotFound) {
				return nil
			}
			return err2
		}

		return item.Value(func(val []byte) error {
			if val == nil {
				return nil
			}
			return obj.Unmarshal(val)
		})
	})
}

func buildKey(ctx context.Context, key []byte, inNamespace ...bool) ([]byte, error) {
	needNamespace := true
	if len(inNamespace) > 0 {
		needNamespace = inNamespace[0]
	}
	var keys [][]byte
	if needNamespace {
		ns, err := ns.NamespaceMust(ctx)
		if err != nil {
			return []byte("."), err
		}
		keys = [][]byte{keyVersion, keyNamespaces, []byte(ns)}
	} else {
		keys = [][]byte{keyVersion}
	}

	b := bytes.Join(keys, keySep)
	return append(b, key...), nil
}

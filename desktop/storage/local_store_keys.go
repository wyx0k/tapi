package storage

import (
	"bytes"
	"errors"
	"strconv"
)

/*
*
SchemaVersion v1

v1
|___session
	|___tab_session
	|___
|___namespaces
	|___{ns_id}
		|___ns_id
		|___ns_name
		|___create_at
		|___update_at
		|___owner
		|___projects
			|___{project_id}
				|___pj_id
				|___pj_name
				|___create_at
				|___update_at
				|___owner
				|___collections
				|	|___{collection_id}
				|		|___create_at
				|		|___update_at
				|		|___requests
				|			|___{request_id}
				|				|___create_at
				|				|___update_at
				|___environments
					|___{env_id}
						|___env_name
						|___create_at
						|___update_at
						|___script
*/

const SchemaVersion = "v1"

var (
	keyVersion = []byte(SchemaVersion)
	keySep     = []byte(".")

	keyNamespaces   = []byte("namespaces")
	keyProjects     = []byte("projects")
	keyCollections  = []byte("collections")
	keyEnvironments = []byte("environments")
	keyRequests     = []byte("requests")
	keySession      = []byte("session")
)

var (
	ErrKeyNotFound    = errors.New("key not found")
	ErrWrongTypeForId = errors.New("wrong type for id")
)

type StoreKey []byte

func (s StoreKey) IsSubKey() bool {
	return bytes.HasPrefix(s, keySep)
}

func (s StoreKey) getType() []byte {
	items := bytes.Split(s, keySep)
	if len(items) < 2 {
		return []byte("")
	}
	last := items[len(items)-1]
	if bytes.Equal(last, keySession) {
		return keySession
	}
	btype := items[len(items)-2]
	switch string(btype) {
	case string(keyNamespaces), string(keyProjects), string(keyCollections), string(keyEnvironments), string(keyRequests):
		return btype
	default:
		return []byte("")
	}
}

func (s StoreKey) GetId() (uint64, error) {
	t := s.getType()
	switch string(t) {
	case string(keyProjects), string(keyCollections), string(keyEnvironments), string(keyRequests):
		items := bytes.Split(s, keySep)
		last := items[len(items)-1]
		return strconv.ParseUint(string(last), 10, 64)
	default:
		return 0, ErrWrongTypeForId
	}
}

func (s StoreKey) GetNamespace() (string, error) {
	if s.IsSubKey() {
		return "", ErrKeyNotFound
	}
	items := bytes.Split(s, keySep)
	ns := items[2]
	return string(ns), nil
}

func Uint64ToBytes(i uint64) []byte {
	s := strconv.FormatUint(i, 10)
	return []byte(s)
}

func BuildBytesSubKey(keys ...[]byte) []byte {
	b := bytes.Join(keys, keySep)
	return append(keySep, b...)
}

func Request(project uint64, collection uint64, request uint64) StoreKey {
	keys := [][]byte{keyProjects, Uint64ToBytes(project), keyCollections, Uint64ToBytes(collection), keyRequests, Uint64ToBytes(request)}
	return BuildBytesSubKey(keys...)
}

func Collection(project uint64, collection uint64) StoreKey {
	keys := [][]byte{keyProjects, Uint64ToBytes(project), keyCollections, Uint64ToBytes(collection)}
	return BuildBytesSubKey(keys...)
}

func Environment(project uint64, environment uint64) StoreKey {
	keys := [][]byte{keyProjects, Uint64ToBytes(project), keyEnvironments, Uint64ToBytes(environment)}
	return BuildBytesSubKey(keys...)
}

func Project(project uint64) StoreKey {
	keys := [][]byte{keyProjects, Uint64ToBytes(project)}
	return BuildBytesSubKey(keys...)
}

func Session() StoreKey {
	return append(keySep, keySession...)
}

type StoreKeys struct {
	Type []byte
}

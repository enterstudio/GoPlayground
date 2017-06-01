package storage

import (
	"errors"
	"sync"
)

type inMemStore struct {
	db map[string][]byte
	m  sync.RWMutex
}

type DB interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte) error
}

var ErrNotFound = errors.New("Error Not Found")

func NewInMemoryDB() DB {
	return &inMemStore{db: make(map[string][]byte)}
}

func (d *inMemStore) Get(key string) ([]byte, error) {
	d.m.Lock()
	defer d.m.Unlock()
	v, ok := d.db[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

func (d *inMemStore) Set(key string, val []byte) error {
	d.m.Lock()
	defer d.m.Unlock()
	d.db[key] = val
	return nil
}

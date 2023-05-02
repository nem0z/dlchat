package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Store struct {
	DB *leveldb.DB
}

func Init(path string) (*Store, error) {
	db, err := leveldb.OpenFile(path, nil)
	return &Store{db}, err
}

func (s *Store) Close() {
	s.DB.Close()
}

func (s *Store) Get(key []byte) ([]byte, error) {
	return s.DB.Get(key, nil)
}

func (s *Store) Put(key []byte, value []byte) error {
	return s.DB.Put(key, value, nil)
}

func (s *Store) Delete(key []byte) error {
	return s.DB.Delete(key, nil)
}

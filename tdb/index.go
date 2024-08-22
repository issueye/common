package tdb

import (
	badgerhold "github.com/timshannon/badgerhold/v4"
)

func NewTDB(path string) (*badgerhold.Store, error) {
	options := badgerhold.DefaultOptions
	options.Dir = path
	options.ValueDir = path

	store, err := badgerhold.Open(options)
	if err != nil {
		return nil, err
	}

	return store, nil
}

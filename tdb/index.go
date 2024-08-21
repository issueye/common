package tdb

import (
	"github.com/dgraph-io/badger/v4"
	badgerhold "github.com/timshannon/badgerhold/v4"
)

type TDB struct {
	*badgerhold.Store
}

func NewTDB(path string) (*TDB, error) {
	opts := badger.DefaultOptions(path)
	store, err := badgerhold.Open(badgerhold.Options{
		Options: opts,
	})
	if err != nil {
		return nil, err
	}

	return &TDB{
		Store: store,
	}, nil
}

func (tdb *TDB) Close() error {
	return tdb.Store.Close()
}

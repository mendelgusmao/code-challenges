package database

import (
	"github.com/mendelgusmao/tax-challenge/backend/config"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func Instance() *bbolt.DB {
	return db
}

func openDatabase(backend config.Specification) error {
	var err error
	db, err = bbolt.Open(config.Backend.Database, 0600, nil)

	if err != nil {
		return errors.Wrap(err, "opening database")
	}

	return nil
}

func init() {
	config.AfterLoad(openDatabase)
}

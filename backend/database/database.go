package database

import (
	"log"
	"sync"

	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/config"
	"go.etcd.io/bbolt"
)

var (
	db   *bbolt.DB
	once sync.Once
)

func Instance() *bbolt.DB {
	once.Do(func() {
		var err error
		db, err = bbolt.Open(config.Backend.Database, 0600, nil)

		if err != nil {
			log.Printf("opening database: %v", err)
		}
	})

	return db
}

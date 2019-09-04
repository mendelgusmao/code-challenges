package services

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

type ProductStorage struct {
	db     *bbolt.DB
	bucket string
}

func NewProductStorage(db *bbolt.DB) *ProductStorage {
	return &ProductStorage{
		db:     db,
		bucket: "products",
	}
}

func (ts *ProductStorage) Store(product TaxedProduct) (uint64, error) {
	var id uint64

	err := ts.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(ts.bucket))

		if err != nil {
			return err
		}

		id, _ = bucket.NextSequence()
		buffer := bytes.NewBuffer([]byte{})

		if err := json.NewEncoder(buffer).Encode(product); err != nil {
			return err
		}

		idBuffer := make([]byte, 8)
		binary.BigEndian.PutUint64(idBuffer, uint64(id))

		return bucket.Put(idBuffer, buffer.Bytes())
	})

	return id, errors.Wrap(err, "ProductStorage.Store")
}

func (ts *ProductStorage) Retrieve(id uint64) (TaxedProduct, error) {
	var product TaxedProduct

	err := ts.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(ts.bucket))

		if err != nil {
			return err
		}

		idBuffer := make([]byte, 8)
		binary.BigEndian.PutUint64(idBuffer, uint64(id))

		content := bucket.Get(idBuffer)
		buffer := bytes.NewBuffer(content)

		return json.NewDecoder(buffer).Decode(&product)
	})

	return product, err
}

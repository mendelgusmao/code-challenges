package services

import (
	"io/ioutil"
	"os"
	"testing"

	"go.etcd.io/bbolt"
)

func TestProductStorageStore(t *testing.T) {
	tests := []struct {
		product       TaxedProduct
		expectedID    uint64
		expectedError error
	}{
		{
			product: TaxedProduct{
				Product: Product{
					Name: "dummy",
				},
			},
			expectedID: 1,
		},
	}

	for _, test := range tests {
		tempdb, err := ioutil.TempFile("/tmp", "products.*.boltdb")
		defer os.Remove(tempdb.Name())

		if err != nil {
			t.Fatal(err)
		}

		db, err := bbolt.Open(tempdb.Name(), 0600, nil)

		if err != nil {
			t.Fatal(err)
		}

		storage := NewProductStorage(db)
		id, err := storage.Store(test.product)

		if err != test.expectedError {
			t.Fatalf("expected error %v, got %v", test.expectedError, err)
		}

		if id != test.expectedID {
			t.Fatalf("expected id %v, got %v", test.expectedID, err)
		}
	}
}

func TestProductStorageRetrieve(t *testing.T) {
	tests := []struct {
		product       TaxedProduct
		expectedError error
	}{
		{
			product: TaxedProduct{
				Product: Product{
					Name: "dummy",
				},
			},
		},
	}

	for _, test := range tests {
		tempdb, err := ioutil.TempFile("/tmp", "products.*.boltdb")
		defer os.Remove(tempdb.Name())

		if err != nil {
			t.Fatal(err)
		}

		db, err := bbolt.Open(tempdb.Name(), 0600, nil)

		if err != nil {
			t.Fatal(err)
		}

		storage := NewProductStorage(db)
		id, err := storage.Store(test.product)

		if err != nil {
			t.Fatal(err)
		}

		retrievedProduct, err := storage.Retrieve(id)

		if err != test.expectedError {
			t.Fatalf("expected error %v, got %v", test.expectedError, err)
		}

		if retrievedProduct != test.product {
			t.Fatalf("expected product %v, got %v", test.product, retrievedProduct)
		}
	}
}

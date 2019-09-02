package services

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

const carsBucket = "cars"

type Car struct {
	ID    int `json:"id"`
	Seats int `json:"seats"`
}

type CarsService struct {
	db *bbolt.DB
}

func NewCarsService(db *bbolt.DB) *CarsService {
	return &CarsService{
		db: db,
	}
}

func (s *CarsService) Clear() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if err := tx.DeleteBucket([]byte(carsBucket)); err != nil {
			return errors.Wrap(err, "clearing cars")
		}

		return nil
	})
}

func (s *CarsService) Put(cars []Car) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(carsBucket))

		if err != nil {
			return errors.Wrap(err, "putting cars")
		}

		for _, car := range cars {
			buffer := bytes.Buffer{}
			enc := gob.NewEncoder(&buffer)

			if err := enc.Encode(car); err != nil {
				return errors.Wrapf(err, "putting cars: encoding car#%d", car.ID)
			}

			err := bucket.Put(
				[]byte(fmt.Sprintf("%d", car.ID)),
				buffer.Bytes(),
			)

			if err != nil {
				return errors.Wrapf(err, "putting cars: putting car#%d", car.ID)
			}
		}

		return nil
	})
}

func (s *CarsService) Find(carID int) (Car, error) {
	var car Car

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(carsBucket))

		carData := bucket.Get([]byte(fmt.Sprintf("%d", carID)))

		if carData == nil {
			return errors.Wrapf(errNotFound, "car#%d", carID)
		}

		buffer := bytes.NewBuffer(carData)
		decoder := gob.NewDecoder(buffer)

		if err := decoder.Decode(&car); err != nil {
			return errors.Wrapf(err, "finding car: decoding car#%d", carID)
		}

		return nil
	})

	if err != nil {
		return car, err
	}

	return car, nil
}

func (s *CarsService) All() ([]Car, error) {
	var cars []Car

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(tripBucket))

		cursor := bucket.Cursor()

		for carID, carData := cursor.First(); carID != nil; carID, carData = cursor.Next() {
			var car Car
			buffer := bytes.NewBuffer(carData)
			decoder := gob.NewDecoder(buffer)

			if err := decoder.Decode(&car); err != nil {
				return errors.Wrapf(err, "fetching cars: decoding car#%d", carID)
			}

			cars = append(cars, car)
		}

		return nil
	})

	if err != nil {
		return cars, err
	}

	return cars, nil
}

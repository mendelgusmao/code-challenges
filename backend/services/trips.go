package services

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

type Trip struct {
	ID        int `json:"id"`
	JourneyID int `json:"journey_id"`
	CarID     int `json:"car_id"`
}

type TripsService struct {
	db     *bbolt.DB
	bucket []byte
}

func NewTripsService(db *bbolt.DB) *TripsService {
	return &TripsService{
		db:     db,
		bucket: []byte("trips"),
	}
}

func (s *TripsService) Clear() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if err := tx.DeleteBucket([]byte(s.bucket)); err != nil {
			return errors.Wrap(err, "clearing trip")
		}

		return nil
	})
}

func (s *TripsService) Insert(trip Trip) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(s.bucket))

		if err != nil {
			return errors.Wrap(err, "putting trip")
		}

		buffer := bytes.Buffer{}
		encoder := gob.NewEncoder(&buffer)

		if err := encoder.Encode(trip); err != nil {
			return errors.Wrapf(err, "inserting trip: encoding trip#%d", trip.ID)
		}

		insertErr := bucket.Put(
			[]byte(fmt.Sprintf("%d", trip.ID)),
			buffer.Bytes(),
		)

		if insertErr != nil {
			return errors.Wrapf(err, "inserting trip: putting trip#%d", trip.ID)
		}

		return nil
	})
}

func (s *TripsService) FindByJourneyID(journeyID int) (Trip, error) {
	var trip Trip

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(s.bucket))

		if bucket == nil {
			log.Printf("bucket '%s' doesnt exist yet", s.bucket)
			return nil
		}

		cursor := bucket.Cursor()

		for tripID, tripData := cursor.First(); tripID != nil; tripID, tripData = cursor.Next() {
			buffer := bytes.NewBuffer(tripData)
			decoder := gob.NewDecoder(buffer)

			if err := decoder.Decode(&trip); err != nil {
				return errors.Wrapf(err, "finding trip: decoding trip#%d", tripID)
			}

			if trip.JourneyID == journeyID {
				break
			}
		}

		return nil
	})

	if err != nil {
		return trip, err
	}

	return trip, nil
}

func (s *TripsService) FindByCarID(carID int) ([]Trip, error) {
	var trips []Trip

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(s.bucket))

		cursor := bucket.Cursor()

		for tripID, tripData := cursor.First(); tripID != nil; tripID, tripData = cursor.Next() {
			buffer := bytes.NewBuffer(tripData)
			decoder := gob.NewDecoder(buffer)
			var trip Trip

			if err := decoder.Decode(&trip); err != nil {
				return errors.Wrapf(err, "finding trip: decoding trip#%d", tripID)
			}

			if trip.CarID == carID {
				trips = append(trips, trip)
			}
		}

		return nil
	})

	if err != nil {
		return trips, err
	}

	return trips, nil
}

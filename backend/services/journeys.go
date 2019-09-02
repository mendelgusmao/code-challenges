package services

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

var errNotFound = errors.New("object not found")

type Journey struct {
	ID     int `json:"id"`
	People int `json:"people"`
}

type JourneysService struct {
	db     *bbolt.DB
	bucket []byte
}

func NewJourneysService(db *bbolt.DB) *JourneysService {
	return &JourneysService{
		db:     db,
		bucket: []byte("journeys"),
	}
}

func (s *JourneysService) Clear() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if err := tx.DeleteBucket([]byte(s.bucket)); err != nil {
			return errors.Wrap(err, "clearing journey")
		}

		return nil
	})
}

func (s *JourneysService) Insert(journey Journey) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(s.bucket))

		if err != nil {
			return errors.Wrap(err, "putting journey")
		}

		buffer := bytes.Buffer{}
		encoder := gob.NewEncoder(&buffer)

		if err := encoder.Encode(journey); err != nil {
			return errors.Wrapf(err, "inserting journey: encoding journey#%d", journey.ID)
		}

		insertErr := bucket.Put(
			[]byte(fmt.Sprintf("%d", journey.ID)),
			buffer.Bytes(),
		)

		if insertErr != nil {
			return errors.Wrapf(err, "inserting journey: putting journey#%d", journey.ID)
		}

		return nil
	})
}

func (s *JourneysService) Find(journeyID int) (Journey, error) {
	var journey Journey

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(s.bucket))

		if bucket == nil {
			log.Printf("bucket '%s' doesnt exist yet", s.bucket)
			return nil
		}

		journeyData := bucket.Get([]byte(fmt.Sprintf("%d", journeyID)))

		if journeyData == nil {
			return errors.Wrapf(errNotFound, "journey#%d", journeyID)
		}

		buffer := bytes.NewBuffer(journeyData)
		decoder := gob.NewDecoder(buffer)

		if err := decoder.Decode(&journey); err != nil {
			return errors.Wrapf(err, "finding journey: decoding journey#%d", journeyID)
		}

		return nil
	})

	if err != nil {
		return journey, err
	}

	return journey, nil
}

func (s *JourneysService) All() ([]Journey, error) {
	var journeys []Journey

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(s.bucket))

		if bucket == nil {
			log.Printf("bucket '%s' doesnt exist yet", s.bucket)
			return nil
		}

		cursor := bucket.Cursor()

		for journeyID, journeyData := cursor.First(); journeyID != nil; journeyID, journeyData = cursor.Next() {
			var journey Journey
			buffer := bytes.NewBuffer(journeyData)
			decoder := gob.NewDecoder(buffer)

			if err := decoder.Decode(&journey); err != nil {
				return errors.Wrapf(err, "fetching journeys: decoding journey#%d", journeyID)
			}

			journeys = append(journeys, journey)
		}

		return nil
	})

	if err != nil {
		return journeys, err
	}

	return journeys, nil
}

package services

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

const journeyBucket = "journeys"

var errNotFound = errors.New("object not found")

type Journey struct {
	ID     int `json:"id"`
	People int `json:"people"`
}

type journeysService struct {
	db *bbolt.DB
}

func NewJourneysService(db *bbolt.DB) *journeysService {
	return &journeysService{
		db: db,
	}
}

func (s *journeysService) Clear() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if err := tx.DeleteBucket([]byte(journeyBucket)); err != nil {
			return errors.Wrap(err, "clearing journey")
		}

		return nil
	})
}

func (s *journeysService) Insert(journey Journey) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(journeyBucket))

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

func (s *journeysService) Find(journeyID int) (Journey, error) {
	var journey Journey

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(journeyBucket))

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

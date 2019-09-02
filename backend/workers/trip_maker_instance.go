package workers

import (
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/config"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/database"
)

var TripMaker *tripMaker

func init() {
	config.AfterLoad(func(backend config.Specification) error {
		db := database.Instance()
		TripMaker = newTripMaker(db, backend.TripMakerInterval)

		return nil
	})
}

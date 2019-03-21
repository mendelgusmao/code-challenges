package partners

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const earthRadius = 6371.0

type dao struct {
	db *mongo.Database
}

func newDAO(db *mongo.Database) *dao {
	return &dao{db: db}
}

func (d *dao) findByLocationAndService(service string, lat, long float64, distance float64) ([]Partner, error) {
	radianDistance := distance / earthRadius

	return d.find(bson.M{
		"availableServices": bson.M{
			"$all": bson.A{service},
		},
		"location.geo": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": bson.A{
					bson.A{lat, long},
					radianDistance,
				},
			},
		},
	})
}

func (d *dao) find(filter interface{}) ([]Partner, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := d.db.Collection("partners").Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	partners := make([]Partner, 0)

	for cursor.Next(ctx) {
		var partner Partner

		err := cursor.Decode(&partner)

		if err != nil {
			return nil, err
		}

		partners = append(partners, partner)
	}

	return partners, nil
}

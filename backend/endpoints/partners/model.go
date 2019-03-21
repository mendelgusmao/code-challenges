package partners

import "go.mongodb.org/mongo-driver/bson/primitive"

type Partner struct {
	ObjectID          *primitive.ObjectID `bson:"_id" json:"-"`
	ID                string              `bson:"id" json:"id"`
	Name              string              `bson:"name" json:"name"`
	AvailableServices []string            `bson:"availableServices" json:"availableServices"`
	Location          struct {
		Name    string  `bson:"name" json:"name"`
		Address string  `bson:"address" json:"address"`
		City    string  `bson:"city" json:"city"`
		State   string  `bson:"state" json:"state"`
		Country string  `bson:"country" json:"country"`
		Lat     float64 `bson:"lat" json:"lat"`
		Long    float64 `bson:"long" json:"long"`
	} `bson:"location" json:"location"`
}

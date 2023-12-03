package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Episode struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	AirDate     time.Time          `bson:"airDate"`
	Description string             `bson:"description"`
	SeasonId    primitive.ObjectID `bson:"seasonId"`
	Season      Season             `json:"season"`
}

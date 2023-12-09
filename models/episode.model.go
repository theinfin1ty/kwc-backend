package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Episode struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	AirDate     time.Time          `bson:"airDate" json:"airDate"`
	Description string             `bson:"description" json:"description"`
	SeasonId    primitive.ObjectID `bson:"seasonId" json:"seasonId"`
	Season      Season             `json:"season"`
}

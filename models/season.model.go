package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Season struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AirDate     time.Time          `bson:"airDate" json:"airDate"`
	Description string             `bson:"description" json:"description"`
}

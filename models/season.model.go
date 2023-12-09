package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Season struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AirDate   time.Time          `bson:"airDate" json:"airDate"`
	Title     string             `bson:"title" json:"title"`
	Subtitle  string             `bson:"subtitle" json:"subtitle"`
	Theme     string             `bson:"theme" json:"theme"`
	Thumbnail string             `bson:"thumbnail" json:"thumbnail"`
	Url       string             `bson:"url" json:"url"`
	Episodes  []Episode          `json:"episodes"`
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Season struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AirDate   time.Time          `bson:"airDate,omitempty" json:"airDate"`
	Title     string             `bson:"title,omitempty" json:"title"`
	Subtitle  string             `bson:"subtitle,omitempty" json:"subtitle"`
	Theme     string             `bson:"theme,omitempty" json:"theme"`
	Thumbnail string             `bson:"thumbnail,omitempty" json:"thumbnail"`
	Url       string             `bson:"url,omitempty" json:"url"`
}

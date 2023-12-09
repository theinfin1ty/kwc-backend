package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Episode struct {
	Id            primitive.ObjectID   `bson:"_id" json:"id"`
	Title         string               `bson:"title" json:"title"`
	AirDate       time.Time            `bson:"airDate" json:"airDate"`
	Subtitle      string               `bson:"subtitle" json:"subtitle"`
	Thumbnail     string               `bson:"thumbnail" json:"thumbnail"`
	Url           string               `bson:"url" json:"url"`
	SeasonId      primitive.ObjectID   `bson:"seasonId" json:"seasonId"`
	Season        Season               `json:"season"`
	ContestantIds []primitive.ObjectID `bson:"contestantIds" json:"contestantIds"`
	Contestants   []Contestant         `json:"contestants"`
	WinnerIds     []primitive.ObjectID `bson:"winnerIds" json:"winnerIds"`
	Winners       []Contestant         `json:"winners"`
}

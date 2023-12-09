package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Episode struct {
	Id            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title         string               `bson:"title,omitempty" json:"title"`
	AirDate       time.Time            `bson:"airDate,omitempty" json:"airDate"`
	Subtitle      string               `bson:"subtitle,omitempty" json:"subtitle"`
	Thumbnail     string               `bson:"thumbnail,omitempty" json:"thumbnail"`
	Url           string               `bson:"url,omitempty" json:"url"`
	SeasonId      primitive.ObjectID   `bson:"seasonId,omitempty" json:"seasonId"`
	Season        []Season             `bson:"season,omitempty" json:"season"`
	ContestantIds []primitive.ObjectID `bson:"contestantIds,omitempty" json:"contestantIds"`
	Contestants   []Contestant         `bson:"contestants,omitempty" json:"contestants"`
	WinnerId      primitive.ObjectID   `bson:"winnerId,omitempty" json:"winnerId"`
	Winner        []Contestant         `bson:"winner,omitempty" json:"winner"`
	RunnerUpId    primitive.ObjectID   `bson:"runnerUpId,omitempty" json:"runnerUpId"`
	RunnerUp      []Contestant         `bson:"runnerUp,omitempty" json:"runnerUp"`
}

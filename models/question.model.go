package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Question  string             `bson:"question,omitempty" json:"question"`
	Answers   []string           `bson:"answers,omitempty" json:"answers"`
	Images    []string           `bson:"images,omitempty" json:"images"`
	Hints     []string           `bson:"hints,omitempty" json:"hints"`
	EpisodeId primitive.ObjectID `bson:"episodeId,omitempty" json:"episodeId"`
	Episode   []Episode          `bson:"episode,omitempty" json:"episode"`
}

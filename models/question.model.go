package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Question  string             `bson:"question" json:"question"`
	Answers   []string           `bson:"answers" json:"answers"`
	Images    []string           `bson:"images" json:"images"`
	Hints     []string           `bson:"hints" json:"hints"`
	EpisodeId primitive.ObjectID `bson:"episodeId" json:"episodeId"`
	Episode   Episode            `json:"episode"`
}

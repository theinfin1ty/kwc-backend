package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Question  string             `bson:"question" json:"question"`
	Answer    string             `bson:"answer" json:"answer"`
	EpisodeId primitive.ObjectID `bson:"episodeId" json:"episodeId"`
	Episode   Episode            `json:"episode"`
}

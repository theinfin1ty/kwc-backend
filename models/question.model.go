package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	Id        primitive.ObjectID `bson:"_id"`
	Question  string             `bson:"question"`
	Answer    string             `bson:"answer"`
	EpisodeId string             `bson:"episodeId"`
	Episode   Episode            `json:"episode"`
}

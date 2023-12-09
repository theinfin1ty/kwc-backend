package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Socials struct {
	Id   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
	Url  string             `bson:"url" json:"url"`
}

type Contestant struct {
	Id      primitive.ObjectID `bson:"_id" json:"id"`
	Name    string             `bson:"name" json:"name"`
	Image   string             `bson:"image" json:"image"`
	Socials []Socials          `bson:"socials" json:"socials"`
}

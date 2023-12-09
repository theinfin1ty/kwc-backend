package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Social struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name,omitempty" json:"name"`
	Url  string             `bson:"url,omitempty" json:"url"`
}

type Contestant struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `bson:"name,omitempty" json:"name"`
	Image   string             `bson:"image,omitempty" json:"image"`
	Socials []Social           `bson:"socials,omitempty" json:"socials"`
}

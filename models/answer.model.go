package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Answer struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId     primitive.ObjectID `bson:"userId" json:"userId"`
	User       User               `json:"user"`
	QuestionId primitive.ObjectID `bson:"questionId" json:"questionId"`
	Question   Question           `json:"question"`
	Answer     string             `bson:"answer" json:"answer"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}

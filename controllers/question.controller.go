package controllers

import (
	"context"
	"kwc-backend/configs"
	"kwc-backend/helpers"
	"kwc-backend/models"
	"kwc-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var QuestionCollection *mongo.Collection = configs.GetCollection(configs.DB, "questions")

func CreateQuestion(c *gin.Context) {
	var body validations.QuestionInput
	var episode models.Episode

	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	episodeId, err := primitive.ObjectIDFromHex(body.EpisodeId)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	err = EpisodeCollection.FindOne(context.TODO(), bson.M{"_id": episodeId}).Decode(&episode)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Episode not found"))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	question := models.Question{
		Id:        primitive.NewObjectID(),
		Question:  body.Question,
		Answers:   body.Answers,
		Images:    body.Images,
		Hints:     body.Hints,
		EpisodeId: episodeId,
	}

	_, err = QuestionCollection.InsertOne(context.TODO(), question)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"question": question}))
	return
}

func ListQuestions(c *gin.Context) {
	questions := []models.Question{}

	episodeId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	cursor, err := QuestionCollection.Aggregate(context.TODO(), []bson.M{
		{"$match": bson.M{"episodeId": episodeId}},
		{"$sort": bson.M{"createdAt": -1}},
		{"$lookup": bson.M{
			"from":         "episodes",
			"localField":   "episodeId",
			"foreignField": "_id",
			"as":           "episode",
		}},
	})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Questions not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	err = cursor.All(context.TODO(), &questions)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"questions": questions}))
	return
}

func GetQuestion(c *gin.Context) {
	var question models.Question

	questionId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	err = QuestionCollection.FindOne(context.TODO(), bson.M{"_id": questionId}).Decode(&question)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Question not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"question": question}))
	return
}

func UpdateQuestion(c *gin.Context) {
	var body validations.QuestionInput
	var question models.Question

	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	questionId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	err = QuestionCollection.FindOne(context.TODO(), bson.M{"_id": questionId}).Decode(&question)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Question not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	question.Question = body.Question
	question.Answers = body.Answers
	question.Images = body.Images
	question.Hints = body.Hints

	_, err = QuestionCollection.UpdateOne(context.TODO(), bson.M{"_id": questionId}, question)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"question": question}))
	return
}

func DeleteQuestion(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid ID"))
		return
	}

	_, err = QuestionCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"message": "Question deleted successfully"}))
	return
}

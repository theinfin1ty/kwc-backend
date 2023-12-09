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

var ContestantCollection *mongo.Collection = configs.GetCollection(configs.DB, "contestants")

func AddContestant(c *gin.Context) {
	var body validations.ContestantInput

	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	var socials []models.Social

	for _, socialInput := range body.Socials {
		socials = append(socials, models.Social{
			Id:   primitive.NewObjectID(),
			Name: socialInput.Name,
			Url:  socialInput.Url,
		})
	}

	contestant := models.Contestant{
		Id:      primitive.NewObjectID(),
		Name:    body.Name,
		Image:   body.Image,
		Socials: socials,
	}

	_, err = ContestantCollection.InsertOne(context.TODO(), contestant)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"contestant": contestant}))
	return
}

func ListContestants(c *gin.Context) {
	var contestants []models.Contestant

	cursor, err := ContestantCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	err = cursor.All(context.TODO(), &contestants)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"contestants": contestants}))
	return
}

func GetContestant(c *gin.Context) {
	var contestant models.Contestant

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid ID"))
		return
	}

	err = ContestantCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&contestant)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"contestant": contestant}))
	return
}

func UpdateContestant(c *gin.Context) {
	var body validations.ContestantInput

	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid ID"))
		return
	}

	var socials []models.Social

	for _, socialInput := range body.Socials {
		socials = append(socials, models.Social{
			Id:   primitive.NewObjectID(),
			Name: socialInput.Name,
			Url:  socialInput.Url,
		})
	}

	contestant := models.Contestant{
		Name:    body.Name,
		Image:   body.Image,
		Socials: socials,
	}

	_, err = ContestantCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": contestant})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Contestant not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"contestant": contestant}))
	return
}

func DeleteContestant(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid ID"))
		return
	}

	_, err = ContestantCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Contestant not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"message": "Contestant deleted successfully"}))
	return

}

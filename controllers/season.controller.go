package controllers

import (
	"context"
	"kwc-backend/configs"
	"kwc-backend/helpers"
	"kwc-backend/models"
	"kwc-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var SeasonCollection *mongo.Collection = configs.GetCollection(configs.DB, "seasons")

func CreateSeason(c *gin.Context) {
	var body validations.SeasonInput

	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	season := models.Season{
		Id:        primitive.NewObjectID(),
		AirDate:   body.AirDate,
		Title:     body.Title,
		Subtitle:  body.Subtitle,
		Theme:     body.Theme,
		Thumbnail: body.Thumbnail,
		Url:       body.Url,
	}

	_, err = SeasonCollection.InsertOne(context.TODO(), season)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"season": season}))
	return
}

func ListSeasons(c *gin.Context) {
	seasons := []models.Season{}

	cursor, err := SeasonCollection.Find(context.TODO(), nil)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Seasons not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	err = cursor.All(context.TODO(), &seasons)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"seasons": seasons}))
	return
}

func GetSeason(c *gin.Context) {
	var season models.Season

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid ID"))
		return
	}

	err = SeasonCollection.FindOne(context.TODO(), models.Season{Id: id}).Decode(&season)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Season not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"season": season}))
	return
}

func UpdateSeason(c *gin.Context) {
	var body validations.SeasonInput

	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return
	}

	season := models.Season{
		Id:        id,
		Title:     body.Title,
		Subtitle:  body.Subtitle,
		Theme:     body.Theme,
		Thumbnail: body.Thumbnail,
		Url:       body.Url,
	}

	_, err = SeasonCollection.UpdateOne(context.TODO(), models.Season{Id: id}, season)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Season not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"season": season}))
	return
}

func DeleteSeason(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return
	}

	_, err = SeasonCollection.DeleteOne(context.TODO(), models.Season{Id: id})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Season not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"message": "Season deleted successfully"}))
	return
}

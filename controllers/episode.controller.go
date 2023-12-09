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

var EpisodeCollection *mongo.Collection = configs.GetCollection(configs.DB, "episodes")

func CreateEpisode(c *gin.Context) {
	var body validations.EpisodeInput
	var season models.Season

	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	seasonId, err := primitive.ObjectIDFromHex(body.SeasonId)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid Season ID"))
		return
	}

	err = SeasonCollection.FindOne(context.TODO(), bson.M{"_id": seasonId}).Decode(&season)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Season not found"))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	var contestantIds []primitive.ObjectID
	var winnerIds []primitive.ObjectID

	for _, cid := range body.ContestantIds {
		contestantId, _ := primitive.ObjectIDFromHex(cid)
		contestantIds = append(contestantIds, contestantId)
	}

	for _, wid := range body.ContestantIds {
		winnerId, _ := primitive.ObjectIDFromHex(wid)
		winnerIds = append(winnerIds, winnerId)
	}

	episode := models.Episode{
		Id:            primitive.NewObjectID(),
		AirDate:       body.AirDate,
		Title:         body.Title,
		Subtitle:      body.Subtitle,
		Thumbnail:     body.Thumbnail,
		Url:           body.Url,
		SeasonId:      seasonId,
		ContestantIds: contestantIds,
		WinnerIds:     winnerIds,
	}

	_, err = EpisodeCollection.InsertOne(context.TODO(), episode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"episode": episode}))
	return
}

func ListEpisodesBySeason(c *gin.Context) {
	episodes := []models.Episode{}

	seasonId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid Season ID"))
		return
	}

	cursor, err := EpisodeCollection.Aggregate(context.TODO(), []bson.M{
		{"$match": bson.M{"seasonId": seasonId}},
		{"$sort": bson.M{"airDate": -1}},
		{"$lookup": bson.M{
			"from":         "seasons",
			"localField":   "seasonId",
			"foreignField": "_id",
			"as":           "season",
		}},
	})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Episodes not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	err = cursor.All(context.TODO(), &episodes)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"episodes": episodes}))
	return
}

func GetEpisode(c *gin.Context) {
	var episode models.Episode

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid ID"))
		return
	}

	err = EpisodeCollection.FindOne(context.TODO(), models.Season{Id: id}).Decode(&episode)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("User not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"episode": episode}))
	return
}

func UpdateEpisode(c *gin.Context) {
	var body validations.EpisodeInput
	var season models.Season

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

	seasonId, err := primitive.ObjectIDFromHex(body.SeasonId)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid Season ID"))
		return
	}

	err = SeasonCollection.FindOne(context.TODO(), bson.M{"_id": seasonId}).Decode(&season)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Season not found"))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	var contestantIds []primitive.ObjectID
	var winnerIds []primitive.ObjectID

	for _, cid := range body.ContestantIds {
		contestantId, _ := primitive.ObjectIDFromHex(cid)
		contestantIds = append(contestantIds, contestantId)
	}

	for _, wid := range body.ContestantIds {
		winnerId, _ := primitive.ObjectIDFromHex(wid)
		winnerIds = append(winnerIds, winnerId)
	}

	episode := models.Episode{
		Id:            id,
		AirDate:       body.AirDate,
		Title:         body.Title,
		Subtitle:      body.Subtitle,
		Thumbnail:     body.Thumbnail,
		Url:           body.Url,
		SeasonId:      seasonId,
		ContestantIds: contestantIds,
		WinnerIds:     winnerIds,
	}

	_, err = SeasonCollection.UpdateOne(context.TODO(), bson.M{"_d": id}, episode)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Episode not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"episode": episode}))
	return
}

func DeleteEpisode(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid ID"))
		return
	}

	_, err = EpisodeCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("Episode not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"message": "Episode deleted successfully"}))
	return
}

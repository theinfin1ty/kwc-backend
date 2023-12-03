package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"kwc-backend/configs"
	"kwc-backend/helpers"
	"kwc-backend/models"
	"kwc-backend/validations"
	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func CreateUser(c *gin.Context) {
	var userInput validations.UserInput
	var user models.User

	err := c.BindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	err = userCollection.FindOne(context.TODO(), bson.M{"email": userInput.Email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No documents found")
		} else {
			c.JSON(http.StatusBadRequest, helpers.InternalServerErrorResponse(err))
			return
		}
	}

	if user != (models.User{}) {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Email already exists"))
		return
	}

	user = models.User{
		Id:        primitive.NewObjectID(),
		Name:      userInput.Name,
		Email:     userInput.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	firebaseAdminAuth, err := configs.FirebaseAdmin.Auth(context.TODO())

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	params := (&auth.UserToCreate{}).Email(userInput.Email).EmailVerified(true).Password(userInput.Password).DisplayName(userInput.Name).Disabled(false)

	firebaseUser, err := firebaseAdminAuth.CreateUser(context.TODO(), params)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	user.Uid = firebaseUser.UID

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, helpers.SuccessResponse(&gin.H{"user": user}))
	return
}

func ListUsers(c *gin.Context) {
	var users []models.User

	results, err := userCollection.Find(context.TODO(), bson.M{}, options.Find().SetProjection(bson.D{{"email", 0}}))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	err = results.All(context.TODO(), &users)

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"users": users}))
	return
}

func GetUser(c *gin.Context) {
	var user models.User

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid ID"))
		return
	}

	err = userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("User not found"))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"user": user}))
	return
}

func UpdateUser(c *gin.Context) {
	var userInput validations.UserInput
	var user models.User

	err := c.BindJSON(&userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid ID"))
		return
	}

	err = userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("User not found"))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":      userInput.Name,
			"email":     userInput.Email,
			"updatedAt": time.Now().UTC(),
		},
	}

	_, err = userCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	if userInput.Password != "" {
		firebaseAdminAuth, err := configs.FirebaseAdmin.Auth(context.TODO())

		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
			return
		}

		params := (&auth.UserToUpdate{}).Password(userInput.Password)

		_, err = firebaseAdminAuth.UpdateUser(context.TODO(), user.Uid, params)

		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
			return
		}
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{}))
	return
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Invalid ID"))
	}

	err = userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	_, err = userCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	firebaseAdminAuth, err := configs.FirebaseAdmin.Auth(context.TODO())

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	err = firebaseAdminAuth.DeleteUser(context.TODO(), user.Uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"user": user}))
}

func GetUserToken(c *gin.Context) {
	var user models.User
	var userInput validations.UserInput

	err := c.BindJSON(&userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequestResponse("Validation Failed"))
		return
	}

	err = userCollection.FindOne(context.TODO(), bson.M{"email": userInput.Email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, helpers.NotFoundResponse("User not found"))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	firebaseAdminAuth, err := configs.FirebaseAdmin.Auth(context.TODO())

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	customToken, err := firebaseAdminAuth.CustomToken(context.TODO(), user.Uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	apiUrl := fmt.Sprintf("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s", configs.GetEnvVariable("FIREBASE_API_KEY"))

	body := struct {
		Token             string `json:"token"`
		ReturnSecureToken bool   `json:"returnSecureToken"`
	}{
		Token:             customToken,
		ReturnSecureToken: true,
	}

	bodyJson, err := json.Marshal(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	response, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(bodyJson))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.InternalServerErrorResponse(err))
		return
	}

	type ResBody struct {
		IdToken string `json:"idToken"`
	}

	var resBody ResBody

	err = json.NewDecoder(response.Body).Decode(&resBody)
	c.JSON(http.StatusOK, helpers.SuccessResponse(&gin.H{"token": resBody}))
	return
}

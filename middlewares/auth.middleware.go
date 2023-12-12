package middlewares

import (
	"context"
	"fmt"
	"kwc-backend/configs"
	"kwc-backend/controllers"
	"kwc-backend/models"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Auth(payload []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		token := c.GetHeader("Authorization")

		token = strings.Split(token, " ")[1]

		firebaseAdminAuth, err := configs.FirebaseAdmin.Auth(context.TODO())

		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		decodedToken, err := firebaseAdminAuth.VerifyIDToken(context.TODO(), token)

		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		err = controllers.UserCollection.FindOne(context.TODO(), bson.M{"uid": decodedToken.UID}).Decode(&user)
		fmt.Println(strings.Split(c.Request.URL.Path, "/"))
		if err != nil {
			pathValues := strings.Split(c.Request.URL.Path, "/")
			if !(err == mongo.ErrNoDocuments && c.Request.Method == "POST" && pathValues[len(pathValues)-1] == "users") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				return
			}
		}

		if len(payload) > 0 && !slices.Contains(payload, user.Role) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		c.Set("user", user)
	}
}

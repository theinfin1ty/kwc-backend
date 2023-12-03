package middlewares

import (
	"context"
	"kwc-backend/configs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(payload []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		token = strings.Split(token, " ")[1]

		firebaseAdminAuth, err := configs.FirebaseAdmin.Auth(context.TODO())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		decodedToken, err := firebaseAdminAuth.VerifyIDToken(context.TODO(), token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.Set("uid", decodedToken.UID)
	}
}

package middleware

import (
	"be-tesis/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		isValid := validateToken(tokenString)
		if !isValid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func validateToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.SecretKey), nil
	})

	if err != nil || !token.Valid {

		return false
	}
	claims := token.Claims.(*jwt.MapClaims)
	if (*claims)["logout"] == true {
		return false
	}

	return true
}

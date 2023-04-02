package middleware

import (
	"myGin/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := services.JWTAuthService().ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claim := token.Claims.(jwt.MapClaims)

		c.Set("user_id", claim["id"])
		c.Next()
	}
}

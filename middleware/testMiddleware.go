package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("From Middleware Test")
		c.Next()
	}
}

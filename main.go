package main

import (
	"myGin/db"
	"myGin/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.DatabaseInit()

	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	api := route.Group("/api")
	routes.SetUserRoutes(api, db)
	route.Run()
}

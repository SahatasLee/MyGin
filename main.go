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
	routes.SetBookRoutes(api, db)
	route.Run(":8000") // set port defualt 8080 -> 8000
}

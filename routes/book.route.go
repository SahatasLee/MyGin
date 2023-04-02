package routes

import (
	"myGin/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetBookRoutes(router *gin.RouterGroup, db *gorm.DB) {
	ctls := controllers.DBController{Database: db}
	router.POST("/book", ctls.CreateBook)
	router.GET("/book/:id", ctls.GetBookById)
	router.GET("/book/list", ctls.GetBookLists)
	router.PATCH("/book", ctls.UpdateBook)
	router.DELETE("/book/:id", ctls.DeleteBookById)
}

package routes

import (
	"myGin/controllers"
	"myGin/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	ctls := controllers.DBController{Database: db}
	router.POST("/register", ctls.Register)
	router.GET("/user/:id", middleware.AuthorizeJWT(), ctls.GetUserById)
	router.GET("/user/list", middleware.AuthorizeJWT(), middleware.Test(), ctls.GetUserLists)
	router.PATCH("/user", ctls.UpdateUser)
	router.DELETE("/user/:id", middleware.AuthorizeJWT(), ctls.DeleteUserById)
	router.POST("/login", ctls.Login)
}

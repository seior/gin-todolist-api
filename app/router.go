package app

import (
	"github.com/gin-gonic/gin"
	"todolist-api/controller"
	"todolist-api/middleware"
)

func NewRouter(controller controller.TodolistController) *gin.Engine {
	router := gin.Default()

	// authorize
	router.Use(middleware.AuthMiddleware())

	todolist := router.Group("/api/v1/todolist")
	{
		todolist.Use(middleware.NotFoundMiddleware())
		todolist.POST("/", controller.Create)
		todolist.PUT("/", controller.Update)
		todolist.DELETE("/:id", controller.DeleteById)
		todolist.GET("/:id", controller.FindById)
		todolist.GET("/", controller.FindAll)
	}

	router.NoRoute(middleware.NotFoundPath)

	return router
}

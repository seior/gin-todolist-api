//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"todolist-api/app"
	"todolist-api/controller"
	"todolist-api/repository"
	"todolist-api/service"
)

func InitializeServer() *gin.Engine {
	wire.Build(app.NewDB, repository.NewTodolistRepository, service.NewTodolistService, controller.NewTodolistController, app.NewRouter)

	return nil
}

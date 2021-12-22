package test

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
	"todolist-api/app"
	"todolist-api/model/request"
	"todolist-api/repository"
	"todolist-api/service"
)

func TestMain(m *testing.M) {
	todolistRepository := repository.NewTodolistRepository()
	db := app.NewDB()
	todolistService := service.NewTodolistService(todolistRepository, db)

	ctx := context.Background()
	todolistService.DeleteAll(ctx)

	m.Run()
}

func TestCreateTodolist(t *testing.T) {
	todolistRepository := repository.NewTodolistRepository()
	db := app.NewDB()
	todolistService := service.NewTodolistService(todolistRepository, db)

	ctx := context.Background()
	todolistRequest := request.CreateTodolistRequest{Name: "maem", Author: "Raden"}
	response := todolistService.Create(ctx, todolistRequest)

	assert.Equal(t, todolistRequest.Name, response.Name)
	assert.Equal(t, todolistRequest.Author, response.Author)
	assert.NotNil(t, response.Id)
}

func TestUpdateTodolistSuccess(t *testing.T) {
	todolistRepository := repository.NewTodolistRepository()
	db := app.NewDB()
	todolistService := service.NewTodolistService(todolistRepository, db)

	ctx := context.Background()
	todolistRequest := request.CreateTodolistRequest{Name: "maem", Author: "Raden"}
	response := todolistService.Create(ctx, todolistRequest)

	updateTodolistRequest := request.UpdateTodolistRequest{
		Id:     response.Id,
		Name:   "Udah Update",
		Author: "Iya udh update",
	}

	todolistService.Update(ctx, updateTodolistRequest)

	assert.NotEqual(t, todolistRequest.Name, updateTodolistRequest.Name)
	assert.NotEqual(t, todolistRequest.Author, updateTodolistRequest.Author)
	assert.NotNil(t, updateTodolistRequest.Id)
}

func TestUpdateTodolistFail(t *testing.T) {
	todolistRepository := repository.NewTodolistRepository()
	db := app.NewDB()
	todolistService := service.NewTodolistService(todolistRepository, db)

	ctx := context.Background()
	todolistRequest := request.CreateTodolistRequest{Name: "maem", Author: "Raden"}
	todolistService.Create(ctx, todolistRequest)

	updateTodolistRequest := request.UpdateTodolistRequest{
		Id:     0,
		Name:   "Udah Update",
		Author: "Iya udh update",
	}

	assert.Panicsf(t, func() {
		todolistService.Update(ctx, updateTodolistRequest)
	}, "todolist not found")
}

func TestDeleteTodolistSuccess(t *testing.T) {
	todolistRepository := repository.NewTodolistRepository()
	db := app.NewDB()
	todolistService := service.NewTodolistService(todolistRepository, db)

	ctx := context.Background()
	todolistRequest := request.CreateTodolistRequest{Name: "maem", Author: "Raden"}
	response := todolistService.Create(ctx, todolistRequest)

	deleteTodolistRequest := request.DeleteTodolistRequest{
		Id: response.Id,
	}

	assert.NotPanics(t, func() {
		todolistService.DeleteById(ctx, deleteTodolistRequest)
	})
}

func TestDeleteTodolisFail(t *testing.T) {
	todolistRepository := repository.NewTodolistRepository()
	db := app.NewDB()
	todolistService := service.NewTodolistService(todolistRepository, db)

	ctx := context.Background()
	todolistRequest := request.CreateTodolistRequest{Name: "maem", Author: "Raden"}
	todolistService.Create(ctx, todolistRequest)

	deleteTodolistRequest := request.DeleteTodolistRequest{
		Id: 0,
	}

	assert.Panicsf(t, func() {
		todolistService.DeleteById(ctx, deleteTodolistRequest)
	}, "todolist not found")
}

func TestFindByIdSuccess(t *testing.T) {
	todolistRepository := repository.NewTodolistRepository()
	db := app.NewDB()
	todolistService := service.NewTodolistService(todolistRepository, db)

	ctx := context.Background()
	todolistRequest := request.CreateTodolistRequest{Name: "maem", Author: "Raden"}
	response := todolistService.Create(ctx, todolistRequest)

	findByIdTodolistRequest := request.FindByIdTodolistRequest{
		Id: response.Id,
	}

	assert.NotPanics(t, func() {
		todolistService.FindById(ctx, findByIdTodolistRequest)
	})
}

func TestFindByIdFail(t *testing.T) {
	todolistRepository := repository.NewTodolistRepository()
	db := app.NewDB()
	todolistService := service.NewTodolistService(todolistRepository, db)

	ctx := context.Background()
	todolistRequest := request.CreateTodolistRequest{Name: "maem", Author: "Raden"}
	todolistService.Create(ctx, todolistRequest)

	findByIdTodolistRequest := request.FindByIdTodolistRequest{
		Id: 0,
	}

	assert.Panicsf(t, func() {
		todolistService.FindById(ctx, findByIdTodolistRequest)
	}, "todolist not found")
}

func TestFindAll(t *testing.T) {
	todolistRepository := repository.NewTodolistRepository()
	db := app.NewDB()
	todolistService := service.NewTodolistService(todolistRepository, db)

	ctx := context.Background()
	todolistRequest := request.CreateTodolistRequest{Name: "maem", Author: "Raden"}

	todolistService.Create(ctx, todolistRequest)
	todolistService.Create(ctx, todolistRequest)
	todolistService.Create(ctx, todolistRequest)
	todolistService.Create(ctx, todolistRequest)
	todolistService.Create(ctx, todolistRequest)

	result := todolistService.FindAll(ctx)

	assert.True(t, len(result) >= 5)
}

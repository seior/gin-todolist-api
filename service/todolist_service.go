package service

import (
	"context"
	"todolist-api/model/request"
	"todolist-api/model/response"
)

type TodolistService interface {
	Create(ctx context.Context, todolistRequest request.CreateTodolistRequest) response.TodolistResponse
	Update(ctx context.Context, todolistRequest request.UpdateTodolistRequest) response.TodolistResponse
	FindById(ctx context.Context, todolistRequest request.FindByIdTodolistRequest) response.TodolistResponse
	FindAll(ctx context.Context) []response.TodolistResponse
	DeleteById(ctx context.Context, todolistRequest request.DeleteTodolistRequest)
	DeleteAll(ctx context.Context)
}

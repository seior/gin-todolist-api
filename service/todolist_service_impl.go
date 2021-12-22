package service

import (
	"context"
	"database/sql"
	"todolist-api/exception"
	"todolist-api/helper"
	"todolist-api/model/request"
	"todolist-api/model/response"
	"todolist-api/repository"
)

type TodolistServiceImpl struct {
	repository.TodolistRepository
	*sql.DB
}

func NewTodolistService(todolistRepository repository.TodolistRepository, DB *sql.DB) TodolistService {
	return &TodolistServiceImpl{TodolistRepository: todolistRepository, DB: DB}
}

func (service *TodolistServiceImpl) Create(ctx context.Context, todolistRequest request.CreateTodolistRequest) response.TodolistResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist := request.ToCreateTodolistRequest(todolistRequest)

	result := service.TodolistRepository.Save(ctx, tx, todolist)

	return response.ToTodolistResponse(result)
}

func (service *TodolistServiceImpl) Update(ctx context.Context, todolistRequest request.UpdateTodolistRequest) response.TodolistResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.TodolistRepository.FindById(ctx, tx, todolistRequest.Id)
	if err != nil {
		panic(exception.NewNotFoundException(err.Error()))
	}

	todolist := request.ToUpdateTodolistRequest(todolistRequest)

	result := service.TodolistRepository.Update(ctx, tx, todolist)

	return response.ToTodolistResponse(result)
}

func (service *TodolistServiceImpl) FindById(ctx context.Context, todolistRequest request.FindByIdTodolistRequest) response.TodolistResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.TodolistRepository.FindById(ctx, tx, todolistRequest.Id)
	if err != nil {
		panic(exception.NewNotFoundException(err.Error()))
	}

	return response.ToTodolistResponse(result)
}

func (service *TodolistServiceImpl) FindAll(ctx context.Context) []response.TodolistResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result := service.TodolistRepository.FindAll(ctx, tx)

	return response.ToTodolistResponses(result)
}

func (service *TodolistServiceImpl) DeleteById(ctx context.Context, todolistRequest request.DeleteTodolistRequest) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.TodolistRepository.FindById(ctx, tx, todolistRequest.Id)
	if err != nil {
		panic(exception.NewNotFoundException(err.Error()))
	}

	service.TodolistRepository.DeleteById(ctx, tx, todolistRequest.Id)
}

func (service *TodolistServiceImpl) DeleteAll(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.TodolistRepository.DeleteAll(ctx, tx)
}

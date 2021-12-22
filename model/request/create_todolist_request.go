package request

import "todolist-api/entity"

type CreateTodolistRequest struct {
	Name   string `json:"name" binding:"required,min=3,max=25" form:"name"`
	Author string `json:"author" binding:"required,min=3,max=255" form:"author"`
}

func ToCreateTodolistRequest(request CreateTodolistRequest) entity.Todolist {
	return entity.Todolist{
		Name:   request.Name,
		Author: request.Author,
	}
}

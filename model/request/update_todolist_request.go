package request

import "todolist-api/entity"

type UpdateTodolistRequest struct {
	Id     int    `json:"id" binding:"required,number" form:"id"`
	Name   string `json:"name" binding:"required,min=3,max=25" form:"name"`
	Author string `json:"author" binding:"required,min=3,max=255" form:"author"`
}

func ToUpdateTodolistRequest(request UpdateTodolistRequest) entity.Todolist {
	return entity.Todolist{
		Id:     request.Id,
		Name:   request.Name,
		Author: request.Author,
	}
}

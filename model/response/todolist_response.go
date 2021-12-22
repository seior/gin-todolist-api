package response

import "todolist-api/entity"

type TodolistResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func ToTodolistResponse(todolist entity.Todolist) TodolistResponse {
	return TodolistResponse{
		Id:     todolist.Id,
		Name:   todolist.Name,
		Author: todolist.Author,
	}
}

func ToTodolistResponses(todolists []entity.Todolist) []TodolistResponse {
	var result []TodolistResponse
	for _, todolist := range todolists {
		response := ToTodolistResponse(todolist)
		result = append(result, response)
	}

	return result
}

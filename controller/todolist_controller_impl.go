package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist-api/model/request"
	"todolist-api/model/response"
	"todolist-api/service"
)

type TodolistControllerImpl struct {
	service.TodolistService
}

func NewTodolistController(todolistService service.TodolistService) TodolistController {
	return &TodolistControllerImpl{TodolistService: todolistService}
}

func (controller *TodolistControllerImpl) Create(c *gin.Context) {
	todolistRequest := request.CreateTodolistRequest{}

	if err := c.ShouldBindJSON(&todolistRequest); err == nil {
		todolist := controller.TodolistService.Create(c.Request.Context(), todolistRequest)

		c.JSONP(http.StatusOK, response.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   todolist,
		})
	} else {
		c.JSONP(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

}

func (controller *TodolistControllerImpl) Update(c *gin.Context) {
	todolistRequest := request.UpdateTodolistRequest{}

	if err := c.ShouldBindJSON(&todolistRequest); err == nil {
		todolist := controller.TodolistService.Update(c.Request.Context(), todolistRequest)

		c.JSONP(http.StatusOK, response.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   todolist,
		})
	} else {
		c.JSONP(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}
}

func (controller *TodolistControllerImpl) DeleteById(c *gin.Context) {
	todolistRequest := request.DeleteTodolistRequest{}

	if err := c.ShouldBindUri(&todolistRequest); err == nil {
		controller.TodolistService.DeleteById(c.Request.Context(), todolistRequest)

		c.JSONP(http.StatusOK, response.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   todolistRequest,
		})
	} else {
		c.JSONP(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}
}

func (controller *TodolistControllerImpl) FindById(c *gin.Context) {
	todolistRequest := request.FindByIdTodolistRequest{}

	if err := c.ShouldBindUri(&todolistRequest); err == nil {
		todolist := controller.TodolistService.FindById(c.Request.Context(), todolistRequest)

		c.JSONP(http.StatusOK, response.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   todolist,
		})
	} else {
		c.JSONP(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}
}

func (controller *TodolistControllerImpl) FindAll(c *gin.Context) {
	todolists := controller.TodolistService.FindAll(c.Request.Context())

	c.JSONP(http.StatusOK, response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todolists,
	})
}

package controller

import "github.com/gin-gonic/gin"

type TodolistController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	DeleteById(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist-api/model/response"
)

func NotFoundPath(c *gin.Context) {
	c.JSONP(http.StatusNotFound, response.WebResponse{
		Code:   http.StatusNotFound,
		Status: "NOT FOUND",
		Data:   "Route not found",
	})
}

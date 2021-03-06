package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist-api/model/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-Key") == "radenganteng" {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Data:   "Wrong api key",
			})
		}
	}
}

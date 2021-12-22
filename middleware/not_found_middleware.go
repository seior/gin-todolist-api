package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist-api/model/response"
)

func NotFoundMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()

			if err != nil {
				c.JSONP(http.StatusNotFound, response.WebResponse{
					Code:   http.StatusNotFound,
					Status: "NOT FOUND",
					Data:   err,
				})
			}
		}()

		c.Next()
	}
}

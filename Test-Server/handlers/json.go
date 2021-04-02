package handlers

import (
	"github.com/gin-gonic/gin"
)

// JSON returns the Handler function for rendering a JSON response
func JSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"hello": "ok",
			"bye":   "not ok",
		})
	}
}

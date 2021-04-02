package handlers

import (
	"github.com/gin-gonic/gin"
)

// Homepage returns the Handler function for rendering the homepage
func Homepage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(200)
		c.File("./static/base.html")
	}
}

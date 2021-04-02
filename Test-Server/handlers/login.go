package handlers

import (
	"github.com/gin-gonic/gin"
)

// LoginPage returns the handler function for rendering the login page
func LoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File("./static/login.html")
	}
}

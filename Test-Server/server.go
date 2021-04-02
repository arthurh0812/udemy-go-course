package main

import (
	// "fmt"
	"github.com/Arthur1208/go-pogramming/Test-Server/handlers"
	// "github.com/Arthur1208/go-pogramming/Test-Server/middleware"
	// "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.Use(static.Serve("/", static.LocalFile("./static", false)))
	router.Static("/static", "./static")

	router.GET("/json", handlers.JSON())

	router.GET("/login", handlers.LoginPage())

	router.GET("/", handlers.Homepage())

	router.Run(":3000")
}

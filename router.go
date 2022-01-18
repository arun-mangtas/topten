package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func initialiseRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("pong")
	})

	return r
}

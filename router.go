package main

import (
	"github.com/gin-gonic/gin"
)

func initialiseRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/top/ten/words", handlerTopTen)

	return r
}

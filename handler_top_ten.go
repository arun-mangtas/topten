package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type handlerTopTenArgs struct {
	Text string `json:"text" binding:"required"`
}

func handlerTopTen(c *gin.Context) {
	var args handlerTopTenArgs
	err := c.BindJSON(&args)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "handlerTopTen",
			"subFunc": "c.BindJSON",
		}).Error(err)
		c.JSON(http.StatusBadRequest, ErrRequestBodyNotProperlyFormatted.Error())
		return
	}

	fmt.Println(args.Text)
	c.Status(http.StatusOK)
	return
}

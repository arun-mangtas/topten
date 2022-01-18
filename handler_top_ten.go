package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mangtas/topten/top"
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

	res, err := top.GetTopResults(args.Text, 10)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "handlerTopTen",
			"subFunc": "top.GetTopResults",
		}).Error(err)
		c.JSON(http.StatusInternalServerError, ErrInternalServerError.Error())
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ramony/gostore/common"
	"github.com/ramony/gostore/dao"
)

func bindDetail(user *gin.RouterGroup) {
	user.GET("/list", listDetails)
}

func listDetails(c *gin.Context) {
	details, err := dao.ListDetails()
	if err != nil {
		response := common.ResponseData{
			Code:    500,
			Status:  "failure",
			Message: "fail to get data",
		}
		c.JSON(400, response)
		return
	}
	response := common.ResponseData{
		Code:   200,
		Status: "success",
		Data:   details,
	}
	c.JSON(200, response)
}

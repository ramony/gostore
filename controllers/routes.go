package controllers

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		detail := v1.Group("/detail")
		bindDetail(detail)
	}
	return router
}

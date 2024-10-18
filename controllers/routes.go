package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFS("/photos", http.Dir("photos"))

	v1 := router.Group("/")
	{
		detail := v1.Group("/detail")
		bindDetail(detail)

		gallery := v1.Group("/gallery")
		bindGallery(gallery)
	}
	return router
}

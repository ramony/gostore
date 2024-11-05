package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFS("/photos", http.Dir("photos"))
	router.StaticFS("/allfiles", http.Dir("allfiles"))

	v1 := router.Group("/")
	{
		detail := v1.Group("/detail")
		bindDetail(detail)

		gallery := v1.Group("/gallery")
		bindGallery(gallery)

		files := v1.Group("/files")
		bindFiles(files)

		bbs := v1.Group("/bbs")
		bindBbs(bbs)

		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index/index.html", gin.H{})
		})
	}
	return router
}

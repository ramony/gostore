package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ramony/gostore/common"
	"github.com/ramony/gostore/dao"
	"github.com/ramony/gostore/models"
)

func bindBbs(bbs *gin.RouterGroup) {
	bbs.GET("/list", listBBS)
	bbs.POST("/save", saveBBS)
}

func saveBBS(c *gin.Context) {
	content := c.Request.FormValue("content")
	if content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "content is required",
		})
		return
	}

	log.Println(content)

	bbsDO := models.BBSDO{Content: content, Status: 0, CreateTime: time.Now()}
	err := dao.SaveBBS(bbsDO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("'%s' save bbs error!", err),
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, "list")
	}
}

func listBBS(c *gin.Context) {
	bbsList, err := dao.ListBBS()
	if err != nil {
		response := common.ResponseData{
			Code:    500,
			Status:  "failure",
			Message: "fail to get data",
		}
		c.JSON(400, response)
		return
	}
	c.HTML(http.StatusOK, "bbs/index.html", gin.H{
		"bbsList": bbsList,
	})

}

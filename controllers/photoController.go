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

func bindGallery(gallery *gin.RouterGroup) {
	gallery.GET("/list", listPhotos)
	gallery.GET("/upload", uploadPage)
	gallery.POST("/upload", uploadPhoto)

}

func uploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{
		"title": "users/index",
	})
}

func uploadPhoto(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println(file.Filename)
	currentTime := time.Now()

	newFileName := fmt.Sprintf("/photos/%d_%s", currentTime.UnixMilli(), file.Filename)
	// 上传文件到指定的目录
	c.SaveUploadedFile(file, "."+newFileName)
	photo := models.Photo{FileName: file.Filename, NewFileName: newFileName, Status: 0}
	err = dao.SavePhoto(photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("'%s' uploaded error!", file.Filename),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded success!", file.Filename),
		})
	}
}

func listPhotos(c *gin.Context) {
	photos, err := dao.ListPhotos()
	if err != nil {
		response := common.ResponseData{
			Code:    500,
			Status:  "failure",
			Message: "fail to get data",
		}
		c.JSON(400, response)
		return
	}
	c.HTML(http.StatusOK, "photos.html", gin.H{
		"photos": photos,
	})

}

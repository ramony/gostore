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

func bindFiles(files *gin.RouterGroup) {
	files.GET("/list", listFiles)
	files.POST("/upload", uploadFiles)
}

func uploadFiles(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println(file.Filename)
	currentTime := time.Now()

	newFileName := fmt.Sprintf("/allfiles/%d_%s", currentTime.UnixMilli(), file.Filename)
	// 上传文件到指定的目录
	c.SaveUploadedFile(file, "."+newFileName)
	fileDO := models.FileDO{FileName: file.Filename, NewFileName: newFileName, Status: 0, CreateTime: time.Now()}
	err = dao.SaveFile(fileDO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("'%s' uploaded error!", file.Filename),
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, "list")
	}
}

func listFiles(c *gin.Context) {
	files, err := dao.ListFiles()
	if err != nil {
		response := common.ResponseData{
			Code:    500,
			Status:  "failure",
			Message: "fail to get data",
		}
		c.JSON(400, response)
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"files": files,
	})

}

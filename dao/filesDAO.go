package dao

import (
	"github.com/ramony/gostore/models"
)

func SaveFile(file models.FileDO) error {
	if err := db.Create(&file).Error; err != nil {
		return nil
	}
	return err
}

func ListFiles() ([]models.FileDO, error) {
	var files []models.FileDO
	if err := db.Order("id desc").Limit(1000).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, err
}

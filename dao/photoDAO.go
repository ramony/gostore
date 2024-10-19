package dao

import (
	"github.com/ramony/gostore/models"
)

func SavePhoto(photo models.Photo) error {
	if err := db.Create(&photo).Error; err != nil {
		return nil
	}
	return err
}

func ListPhotos() ([]models.Photo, error) {
	var photos []models.Photo
	if err := db.Limit(1000).Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, err
}

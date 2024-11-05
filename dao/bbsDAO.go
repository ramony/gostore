package dao

import (
	"github.com/ramony/gostore/models"
)

func SaveBBS(bbs models.BBSDO) error {
	if err := db.Create(&bbs).Error; err != nil {
		return nil
	}
	return err
}

func ListBBS() ([]models.BBSDO, error) {
	var bbsList []models.BBSDO
	if err := db.Order("id desc").Limit(1000).Find(&bbsList).Error; err != nil {
		return nil, err
	}
	return bbsList, err
}

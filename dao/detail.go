package dao

import (
	"github.com/ramony/gostore/models"
)

func ListDetails() ([]models.Detail, error) {
	var details []models.Detail
	if err := db.Limit(6).Find(&details).Error; err != nil {
		return nil, err
	}
	return details, err
}

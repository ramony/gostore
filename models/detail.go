package models

type Detail struct {
	Id         int    `json:"id",gorm:"auto-increment"`
	DetailType string `json:"detail_type"`
}

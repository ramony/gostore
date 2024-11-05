package models

import "time"

type BBSDO struct {
	Id         int       `json:"id",gorm:"auto-increment"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
}

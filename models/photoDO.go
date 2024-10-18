package models

type Photo struct {
	Id          int    `json:"id",gorm:"auto-increment"`
	FileName    string `json:"file_name"`
	NewFileName string `json:"new_file_name"`
	Status      int    `json:"status"`
}

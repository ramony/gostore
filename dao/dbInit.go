package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/ramony/gostore/models"
)

const (
	Mysqldb = "rayuser:123456@tcp(192.168.0.107:3306)/jstoremini?charset=utf8"
)

var db *gorm.DB
var err error

func init() {
	fmt.Println("init db start....")
	db, err = gorm.Open("mysql", Mysqldb)
	if err != nil {
		panic("db error")
	}
	db.SingularTable(true)
	db.AutoMigrate(&models.Photo{})
	//defer DB2.Close()
}

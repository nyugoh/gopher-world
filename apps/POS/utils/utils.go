package qpos

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func DbConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost)/qpos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	if db.DB().Ping() != nil {
		log.Println("Connected successfully to DB...")
	}
	return db
}

package rest

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	table string = "CREATE TABLE IF NOT EXISTS posts ( " +
		"id INT(11) AUTO_INCREMENT," +
		"title VARCHAR(200) NOT NULL," +
		"description VARCHAR(200) NOT NULL," +
		"meta VARCHAR(500) ," +
		"content TEXT NOT NULL ," +
		"image_cover VARCHAR(200) NOT NULL DEFAULT 'cover.png'," +
		"created_at DATETIME ," +
		"modified_at DATETIME ," +
		"deleted_at DATETIME," +
		"PRIMARY KEY(id)," +
		"UNIQUE(title)" +
	")"
)

func CreateTable() {
	db := CreateDBConnection()
	defer db.Close()

	if result, err := db.Exec(table); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(result)
	}
}

func CreateDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(0.0.0.0:3306)/blog_api?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}


package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// Returns a connection to db or error
func DbConnect() (*sql.DB, error) {
	dbUri := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		Log("unable to connect to DB", err.Error())
		return nil, err
	} else {
		Log("Connected to DB")
	}
	return db, nil
}

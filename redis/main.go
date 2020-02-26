package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to read config file:", err.Error())
	}

	client, err := redisConnect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()

	db, err := DbConnect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()


}

func redisConnect() (*redis.Client, error) {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_SERVER"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASS"),
		DB:       db,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to redis:%s", err.Error()))
	}

	fmt.Println("Redis reply: ", pong)
	return client, nil
}

func DbConnect() (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_URI"))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable too connect to DB: %s", err.Error()))
	}

	if db.Ping() != nil {
		return nil, errors.New("unable to ping DB")
	}
	fmt.Println("Connected to DB...")
	return db, nil
}
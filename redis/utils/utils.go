package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func InitLogger() {
	godotenv.Load()
	logFolder := os.Getenv("LOG_FOLDER")
	appName := os.Getenv("APP_NAME")
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s%s-old.%s.log", logFolder, appName, "%Y-%m-%d.%H%M"),
		rotatelogs.WithLinkName(fmt.Sprintf("%s%s.log", logFolder, appName)),
		rotatelogs.WithRotationTime(time.Hour*1),
	)
	if err != nil {
		fmt.Println("Failed to initialize log file ", err.Error())
	}
	log.SetOutput(writer)
	Log("Logger intialized successfully...")
}

func Log(msg ...interface{}) {
	log.Printf("%v", msg)
}

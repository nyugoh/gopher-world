package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func InitLogger() {
	logFolder := os.Getenv("LOG_FOLDER")
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s.log", logFolder+"freebets-old", "%Y-%m-%d.%H%M"),
		rotatelogs.WithLinkName(logFolder+"freebets.log"),
		rotatelogs.WithRotationTime(time.Hour*1),
	)
	if err != nil {
		fmt.Println("Failed to initialize log file ", err.Error())
	}
	log.SetOutput(writer)
	return
}

func Log(msg ...interface{}) {
	fmt.Printf("%s: ", time.Now().String())
	fmt.Print(msg...)
	log.Println(msg...)
}

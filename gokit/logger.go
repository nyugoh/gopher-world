package gokit

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cast"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

var AppName string

// InitLogger Initializes logger, writes log messages to files
// * If you want to log locally (e.g in development) set the env value to either (LOCAL | DEV)
// Log files are named from the APP_NAME and a time stamp attached to rotated files
// * if env is LOCAL, then logs will be output to stdErr
// * this is got from :: https://gitlab.betika.private/betikateam/gokit/-/blob/fff076dec2d179482352f1e513b22c7fc53a3a38/logger.go#L53
// * in case of any difference in opinion, feel free to raise and change
func InitLogger(env, appName, logFolder, logLevel string) error {
	AppName = appName
	if strings.EqualFold(env, "LOCAL") {
		log.SetOutput(os.Stderr)
		return nil
	}
	if strings.EqualFold(env, "DEV") {
		pwd, err := os.Getwd()
		if err == nil {
			logFolder = fmt.Sprintf("%s/logs/", pwd)
		}
	}
	if strings.EqualFold(logFolder, "") {
		return fmt.Errorf("%s log folder is required", LogPrefix)
	}
	if strings.EqualFold(appName, "") {
		return fmt.Errorf("%s app name is required", LogPrefix)
	}

	writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s.json", logFolder+appName+"-old", "%Y-%m-%d"),
		rotatelogs.WithLinkName(logFolder+appName+".json"),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(500),
	)
	if err != nil {
		return fmt.Errorf("%s failed to initialize log file::%s", LogPrefix, err.Error())
	}

	log.SetFormatter(
		&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
			FieldMap: log.FieldMap{
				"app": appName,
			},
		})
	// set output level
	if len(logLevel) == 0 {
		logLevel = "info"
	}
	l, err := log.ParseLevel(logLevel)
	if err != nil {
		fmt.Printf(`%s Bad log level %s: %s\n`, LogPrefix, logLevel, err.Error())
		return err
	}
	log.SetLevel(l)
	log.SetOutput(writer)
	Log("%s Logger initialized successfully", LogPrefix)
	Log("%sLog folder:%s Log level:%v App Name:%s", LogPrefix, logFolder, l, appName)
	return nil
}

func Log(msgFormat string, params ...interface{}) {
	LogInfo(msgFormat, params...)
}

func LogInfo(msgFormat string, params ...interface{}) {
	msg := fmt.Sprintf(msgFormat, params...)
	log.WithField("app", AppName).Info(msg)
}

func LogWarn(msgFormat string, params ...interface{}) {
	msg := fmt.Sprintf(msgFormat, params...)
	log.WithField("app", AppName).Warn(msg)
}

func LogError(msgFormat string, params ...interface{}) {
	msg := fmt.Sprintf(msgFormat, params...)
	log.WithField("app", AppName).Error(msg)
}

func LogDebug(msgFormat string, params ...interface{}) {
	msg := fmt.Sprintf(msgFormat, params...)
	log.WithField("app", AppName).Debug(msg)
}

func msgToJson(msg []interface{}) (s string, err error) {
	switch len(msg) {
	case 0:
		s = "\"\""
	default:
		var castedArgs []string
		for _, b := range msg {
			castedArg, err := cast.ToStringE(b)
			if err != nil {
				fmt.Println(err)
				return "", err
			}
			castedArgs = append(castedArgs, castedArg)
		}
		s = strings.Join(castedArgs, " ")
	}
	return s, nil
}

func removeBraces(msg string) string {
	if strings.HasPrefix(msg, "[") && strings.HasSuffix(msg, "]") {
		msg = msg[1 : len(msg)-1]
		msg = removeBraces(msg)
	}
	if len(msg) <= 1 {
		return msg
	}
	return msg
}

func LogObject(msg string, o interface{})      { LogObjectInfo(msg, o) }
func LogObjectInfo(msg string, o interface{})  { writeObject("INFO", msg, o) }
func LogObjectWarn(msg string, o interface{})  { writeObject("WARN", msg, o) }
func LogObjectError(msg string, o interface{}) { writeObject("ERROR", msg, o) }
func LogObjectDebug(msg string, o interface{}) { writeObject("DEBUG", msg, o) }

func writeObject(logLevel, msg string, obj ...interface{}) {
	objStr, err := json.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("Cannot marshal object: %s", err))
	}
	str := string(objStr)
	if strings.ToUpper(os.Getenv("APP_ENV")) == "LOCAL" {
		fmt.Printf("%v\n", str)
	} else {
		log.Printf(`{"level":"%s", "time":"%s", "app_name":"%s", "message":%q, "data":%s}`, logLevel, CurrentTimeMill(), os.Getenv("APP_NAME"), msg, str)
	}
}

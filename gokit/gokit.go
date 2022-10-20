package gokit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	LogPrefix              = "[gokit]"
	InternalTimeFormat     = "2006-01-02 15:04:05"
	InternalTimeFormatMill = "2006-01-02 15:04:05.000"
)

// CurrentTime Return current time in string format InternalTimeFormat
func CurrentTime() string {
	return time.Now().Format(InternalTimeFormat)
}

// CurrentTimeMill Return current time in string format InternalTimeFormat
func CurrentTimeMill() string {
	return time.Now().Format(InternalTimeFormatMill)
}

// ExitApp - Logs the error and exits app
func ExitApp(err error) {
	LogError(err.Error())
	Log("%s Exiting app...", LogPrefix)
	os.Exit(2) // set to any code != 0
}

func RespondWithError(w http.ResponseWriter, code int, msg string) error {
	return RespondWithJSON(w, code, map[string]interface{}{"response": msg, "status": "failed"})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		return err
	}
	return nil
}

func GetConfigMandatory(param string) (value string) {
	if value, present := os.LookupEnv(param); present {
		return value
	}
	panic(fmt.Sprintf("Environment variable `%s' not found.", param))
}

func GetConfigOptional(param, defaultValue string) (value string) {
	if value, present := os.LookupEnv(param); present {
		return value
	}
	return defaultValue
}

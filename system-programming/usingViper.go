package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("GOMAXPROCS") // bind an env variable to viper
	val := viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)

	viper.Set("GOMAXPROCS", 15)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)

	viper.BindEnv("TEST_ENV")
	val = viper.Get("TEST_ENV")
	if val == nil {
		fmt.Println("TEST_ENV is not defined")
		return
	}
	fmt.Println("TEST_ENV:", val)
}

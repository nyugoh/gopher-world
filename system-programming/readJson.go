package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.SetConfigFile("configs.json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err.Error())
		return
	}

	if viper.IsSet("item1.key1") {
		fmt.Println("item1.key1:", viper.Get("item1.key1"))
	} else {
		fmt.Println("item1.key1 is not set")
	}

	if viper.IsSet("item2.key3") {
		fmt.Println("item2.key3:", viper.Get("item2.key3"))
	} else {
		fmt.Println("item2.key3 is not set")
	}

	if !viper.IsSet("item3.key1") {
		fmt.Println("item3.key1 is not set.")
	}
}

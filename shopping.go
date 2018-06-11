package main

import (
	utils "github.com/nyugoh/go-shopping-cli/utils"
	"fmt"
	"strconv"
)

var (
	menuOption string
)
func main() {
	utils.WelcomeMessage()
	utils.MainMenu()
	fmt.Scan(&menuOption)
	setOption(menuOption)
}

func setOption(s string)  {
	i, _ := strconv.Atoi(s)
	switch i {
	case 0:
		//exit app
		fmt.Println("Exiting app ...")
	case 1:
		//cal
		fmt.Println("selected 1")
	default:
		fmt.Println("No option selected")
	}
}
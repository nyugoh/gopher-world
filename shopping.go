package main

import (
	"github.com/nyugoh/go-shopping-cli/utils"
	"fmt"
	"strconv"
	"github.com/nyugoh/go-shopping-cli/store"
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
		// Add items to store
		store.AddStock()
		fmt.Println(store.Stock)
	default:
		fmt.Println("No option selected")
	}
}
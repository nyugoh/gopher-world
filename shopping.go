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
	for {
		utils.MainMenu()
		fmt.Scan(&menuOption)
		if menuOption == "q" {
			break
		} else {
			setOption(menuOption)
		}
	}
}

func setOption(s string)  {
	i, _ := strconv.Atoi(s)
	switch i {
	case 1:
		store.AddStock()// Add items to store
	case 2:
		store.ListStock() // List all items
	default:
		fmt.Println("No option selected")
	}
}
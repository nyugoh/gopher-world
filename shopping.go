package main

import (
	"fmt"
	"github.com/nyugoh/go-shopping-cli/store"
	"github.com/nyugoh/go-shopping-cli/utils"
	"strconv"
	"github.com/nyugoh/go-shopping-cli/cart"
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

func setOption(s string) {
	i, _ := strconv.Atoi(s) // Convert string input to integer
	switch i {
	case 1:
		store.AddStock() // Add items to store
	case 2:
		store.ListStock() // List all items
	case 3:
		addToCart() // List all items
	default:
		fmt.Println("No option selected")
	}
}

func addToCart()  {
	store.ListStock()
	option := 1
	for {
		if option == 0 {
			break
		}
		itemId := 1
		fmt.Printf("\nEnter item # ::")
		fmt.Scanf("%d", &itemId)
		//TODO: Check if id is greater than total item
		item := store.Stock.Items[itemId]
		var cartItem cart.CartItem
		for _, i := range item {
			cartItem = cart.CartItem{1, i.Name, i.Price}
		}
		cart.AddItem(&cartItem)
		fmt.Println("Item added...")
		fmt.Printf("Select 0. Back 1. Continue 2. Show cart content\n ==>")
		fmt.Scanf("%d", &option)
		if option == 1 {
			continue
		}
		if option == 2 {
			cart.DisplayCart()
		}
	}
}
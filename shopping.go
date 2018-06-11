package main

import (
	"fmt"
	"github.com/nyugoh/go-shopping-cli/cart"
	"github.com/nyugoh/go-shopping-cli/store"
	"github.com/nyugoh/go-shopping-cli/utils"
	"strconv"
	"time"
	"strings"
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
	case 4:
		checkOut() // Display cart items
	default:
		fmt.Println("No option selected")
	}
}

func addToCart() {
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

func checkOut()  {
	option := 1
	cart.DisplayCart()
	fmt.Println("Options 0.Continue shopping 1.Proceed to payment 2.Remove item 3.Empty cart")
	fmt.Scanf("%d", &option)

	switch option {
	case 0:
		addToCart()
	case 1:
		processCart()
	case 2:
		removeFromCart()
	case 3:
		cart.EmptyCart()
	}
}

func processCart()  {
	loading := "-\\|/"
	fmt.Println("Processing cart")
	for i:=0; i<100; i++  {
		for _, symbol := range strings.Split(loading, "") {
			fmt.Printf("\r----------%s----------", symbol)
			time.Sleep(10000*time.Microsecond)
		}
	}
	fmt.Printf("\n Total items ==> %d", len(cart.ShoppingCart.Items))
	fmt.Printf("\n Total price ==> %f\n", cart.ShoppingCart.Total)
}

func removeFromCart()  {
	itemId := 0
	fmt.Printf("Item id to remove ==> ")
	fmt.Scanf("%d", &itemId)
	if itemId > len(cart.ShoppingCart.Items) {
		fmt.Println("Item does't exists in cart")
		return
	}
	cart.RemoveItem(itemId)
}
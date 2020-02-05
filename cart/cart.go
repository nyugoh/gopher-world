package cart

import "fmt"

type CartItem struct {
	Qty   int
	Name  string
	Price float64
}

type Cart struct {
	Items []CartItem
	Total float64
}

var (
	ShoppingCart Cart
)

func AddItem(newItem *CartItem) {
	exists := false
	for idx, item := range ShoppingCart.Items {
		if item.Name == newItem.Name {
			ShoppingCart.Items[idx].Qty += newItem.Qty // Increment qty
			exists = true
		}
	}
	if !exists {
		ShoppingCart.Items = append(ShoppingCart.Items, *newItem)
	}
	ShoppingCart.Total += newItem.Price
}

func DisplayCart() {
	fmt.Println("\n\n\t******* Your shopping cart *******")
	fmt.Println("\n\t----------------------------------")
	fmt.Println("\t| # | Qty | Name | Price | Total |")
	fmt.Println("\t----------------------------------")
	if len(ShoppingCart.Items) == 0 {
		fmt.Println("\tEmpty....!! :) use 3 to add items.")
	}
	for idx, item := range ShoppingCart.Items {
		fmt.Printf("\t|%3d|%5d|%6s|%7.2f|%7.2f|\n", idx, item.Qty, item.Name, item.Price, item.Price * float64(item.Qty))
		fmt.Println("\t----------------------------------")	}
	fmt.Printf("\n\tTotal :: %f\n", ShoppingCart.Total)
	fmt.Println("\t----------------------------------")
}

func RemoveItem(id int)  {
	items := make([]CartItem, 0)
	for idx, item := range ShoppingCart.Items {
		if id == idx {
			ShoppingCart.Total -= item.Price* float64(item.Qty)
			continue
		}
		items = append(items, item)
	}
	ShoppingCart.Items = items
	DisplayCart()
}

func EmptyCart()  {
	ShoppingCart.Total = 0
	ShoppingCart.Items = []CartItem{}
	DisplayCart()
}
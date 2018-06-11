package cart

import "fmt"

type CartItem struct {
	Qty int
	Name string
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
	if len(ShoppingCart.Items) == 0 {
		ShoppingCart.Items = append(ShoppingCart.Items, *newItem)
	} else {
		for idx, item := range ShoppingCart.Items {
			if item.Name == item.Name {
				ShoppingCart.Items[idx].Qty += newItem.Qty // Increment qty
			} else {
				ShoppingCart.Items = append(ShoppingCart.Items, *newItem)
			}
		}
	}
	ShoppingCart.Total += newItem.Price
}

func DisplayCart()  {
	fmt.Println(ShoppingCart)
}

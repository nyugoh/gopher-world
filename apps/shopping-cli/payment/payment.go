package payment

import (
	"../cart"
	"fmt"
)

var (
	Amount   float64
	Total    float64
	Balance  float64
	Discount float64
)

func ProcessOrder() {
	Amount, Total, Balance, Discount = 0.00, cart.ShoppingCart.Total, cart.ShoppingCart.Total, 0.00
	option := "n" // Default to no
	fmt.Println("Apply Discount y/n")
	fmt.Scanf("%s", &option)
	if option == "y" {
		fmt.Print("\nEnter Discount amount ::")
		fmt.Scanf("%f", &Discount)
		Balance -= Discount
	}
	// Display Balance
	for {
		fmt.Printf("\n\nBalance :: %f", Balance)
		fmt.Printf("\nTotal   :: %f", Total)
		amount := 0.00
		fmt.Print("\nAdd payment  ==>")
		fmt.Scanf("%f", &amount)
		Balance -= amount
		Amount += amount
		if Balance == 0.00 {
			break
		}
	}
}

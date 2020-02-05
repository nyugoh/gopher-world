package utils

import (
	"../cart"
	"../payment"
	"fmt"
)

func WelcomeMessage() {
	fmt.Println("\n========= Welcome to the best shopping cart =========")
	fmt.Println(" We offer Gopher products for awesome people like you ")
	fmt.Println("                 Happy shopping")
	fmt.Printf("*****************************************************\n\n\n")
}

func MainMenu() {
	fmt.Println("|------------------ MENU -------------|")
	fmt.Println("              Main Menu")
	fmt.Println("Use the options below to navigate")
	fmt.Println("1. Add item")
	fmt.Println("2. List items")
	fmt.Println("3. Select item")
	fmt.Println("4. Checkout/ Get cart status")
	fmt.Println("5. Pay")
	fmt.Println("6. Print receipt")
	fmt.Println("q - quit the application")
	fmt.Println("|___________________ END ____________|")
	fmt.Printf("\nSelect ==> ")
}

func PrintReceipt() {
	fmt.Printf("\n\t^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	fmt.Println("\t\n     Gopher CLI Shopping")
	fmt.Println("\t   All the geeky stuff you need")
	fmt.Println("\t----------------------------------")
	fmt.Println("\t| # | Qty | Name | Price | Total |")
	fmt.Println("\t----------------------------------")
	for idx, item := range cart.ShoppingCart.Items {
		fmt.Printf("\t|%3d|%5d|%6s|%7.2f|%7.2f|\n", idx, item.Qty, item.Name, item.Price, item.Price*float64(item.Qty))
		fmt.Println("\t----------------------------------")
	}
	fmt.Printf("\n\t|%23s|%7.2f|", "Total", cart.ShoppingCart.Total)
	fmt.Printf("\n\t|%23s|%7.2f|", "Discount", payment.Discount)
	fmt.Printf("\n\t|%23s|%7.2f|", "Balance", payment.Balance)
	fmt.Printf("\n\t|%23s|%7.2f|", "Payments", payment.Amount)
	fmt.Printf("\n\t^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n\n")
}

package utils

import "fmt"

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

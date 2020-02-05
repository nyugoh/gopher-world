package main

import "fmt"

func lcm(a, b int) int {
	var lcmnum, _lcm int
	if a > b {
		lcmnum = b
	} else {
		lcmnum = a
	}
	for {
		if a%lcmnum == 0 && b%lcmnum == 0 {
			_lcm = lcmnum
			break
		}
		lcmnum++
	}

	return _lcm
}

func gcd(a, b int) int {
	var _gcd int

	for idx := 1; idx < a && idx < b; idx++ {
		if a%idx == 0 && b%idx == 0 {
			_gcd = idx
		}
	}
	return _gcd
}

func main() {
	var num1, num2, option int
	fmt.Printf("Enter 2 positive numbers ::\n")
	fmt.Printf("\t--> ")
	fmt.Scanf("%d", &num1)
	fmt.Printf("\t--> ")
	fmt.Scanf("%d", &num2)

	fmt.Println("Select 1 or 2::")
	fmt.Println("\t1. LCM")
	fmt.Println("\t2. GCD")
	fmt.Printf("\t--> ")
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		fmt.Printf("\tLCM:%d\n", lcm(num1, num2))
	case 2:
		fmt.Printf("\tGCD:%d\n", gcd(num1, num2))
	default:
		fmt.Println("\n\nNo option selected.")
	}
}

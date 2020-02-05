package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name, age := "", 0
	fmt.Print("Enter your name ::")
	fmt.Scan(&name)

	fmt.Print("Enter your age ::")
	fmt.Scan(&age)

	fmt.Printf("Your name is %s and you are %d years old.", name, age)

	scanner := bufio.NewReader(os.Stdin)
	fmt.Printf("\nEnter your last name ::")
	name, err := scanner.ReadString('\n')
	fmt.Printf("%s\n", name)
	if err != nil {
		fmt.Println("err", err)
	}
}

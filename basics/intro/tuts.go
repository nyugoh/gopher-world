package intro

import (
	"flag"
	"fmt"
)

func tuts() {
	fmt.Println("Hello world")

	maxFlag := flag.Int("max", 5, "Max value")
	nameArg := flag.String("name", "Jane", "Name of user")
	flag.String("gender", "", "Gender value")
	flag.Parse()

	for i:=0; i < *maxFlag; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	if len(flag.Args()) == 0 {
		fmt.Println("Provide a name")
		return
	}
	fmt.Println("Name==>", *nameArg)
	fmt.Println(flag.Args())
}

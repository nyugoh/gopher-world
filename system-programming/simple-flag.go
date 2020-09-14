package main

import (
	"flag"
	"fmt"
)

func main() {
	minusK := flag.Bool("K", true, "K flag, set debug mode on or off")
	minusO := flag.Int("O", 0, "O flag, default start value")
	flag.Parse()

	*minusO ++
	*minusO ++
	fmt.Println("-K:", *minusK)
	fmt.Println("-O: ", *minusO)
}

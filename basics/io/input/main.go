package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	langs := map[string]int{}
	f, err := os.Create("results.txt")
	if err != nil {
		log.Fatal("Error::", err.Error())
	}
	defer f.Close()
	for scanner.Scan() {
		langs[strings.ToUpper(scanner.Text())]++
	}

	if err := scanner.Err(); err !=nil {
		log.Fatalln("Error::", err.Error())
	}
	for lang, count := range langs {
		f.WriteString(fmt.Sprintf("Lang:: %s ==> %d\n", lang, count))
	}
}

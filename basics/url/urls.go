package main

import (
	"net/url"
	"log"
	"fmt"
)

func main()  {
	urlString := "https://joe:1234@google.com/search?q=dotnet&version=2.34#introduction"
	parsedUrl, err := url.Parse(urlString)

	if err != nil {
		log.Fatal(err)
	}

	qs := parsedUrl.RawQuery
	fmt.Printf("Scheme : %s\n", parsedUrl.Scheme)
	fmt.Printf("User : %s\n", parsedUrl.User)
	fmt.Printf("Host : %s\n", parsedUrl.Host)
	fmt.Printf("Path : %s\n", parsedUrl.Path)
	fmt.Printf("Query params : %s\n", qs)
	parmas, _ := url.ParseQuery(qs)
	fmt.Printf("Parsed query params : %v\n", parmas)
	fmt.Printf("Fragment : %s\n", parsedUrl.Fragment)

}

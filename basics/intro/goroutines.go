package intro

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HomePageSize struct {
	Url string
	Size int
}


func goroutines() {
	resultChannel := make(chan HomePageSize)
	urls := []string {
		"http://localhost:6060",
		"http://localhost:6060/pkg/",
		"http://localhost:6060/pkg/compress/bzip2/",
		"https://google.com",
	}

	for _, url := range urls {
		go parseUrl(url, resultChannel)
	}

	var biggestHomePage HomePageSize
	for range urls {
		res := <- resultChannel
		if res.Size > biggestHomePage.Size {
			biggestHomePage = res
		}
	}
	fmt.Println("Biggest home page is :", biggestHomePage.Url)
}

func parseUrl(url string, c chan<- HomePageSize) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	c <- HomePageSize{
		Url:  url,
		Size: len(bs),
	}
}



func selectChannel() {
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		for {
			ch1 <- "First goroutine"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			ch2 <- "Second goroutine"
			time.Sleep(time.Second * 3)
		}
	}()
	
	go func() {
		for {
			select {
			case msg1 := <- ch1:
				fmt.Println(msg1)
			case msg2:= <- ch2:
				fmt.Println(msg2)
			//case timeElapsed:= <- time.After(time.Second):
			//	fmt.Println("Time elapsed", timeElapsed.Second())
			//default:
			//	fmt.Println("Nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanf("%s", &input)
}

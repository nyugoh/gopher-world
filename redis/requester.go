package main

import (
	"fmt"
	"github.com/nyugoh/gopher-world/redis/utils"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	client *http.Client
	errorChannel chan bool
)
func init()  {
	log.Println("Initializing app...")
	client = &http.Client{}
	errorChannel = make(chan bool)
	utils.InitLogger()
}

func main() {
	startTime := time.Now()
	wg := sync.WaitGroup{}
	for i:=1; i <=1000;i++ {
		wg.Add(1)
		go func() {
			betId := rand.Int63n(1200000)
			msg := fmt.Sprintf("Making request for bet Id:%d", betId)
			utils.Log(msg)
			fmt.Println(msg)
			sleepTime := rand.Intn(5)
			sleep := sleepTime*time.Second
			time.Sleep(time.Second)
			res, err := client.Get(fmt.Sprintf("http://localhost:5000/bets/%d", betId))
			if err != nil {
				utils.Log(fmt.Sprintf("Response error :%s", err.Error()))
				errorChannel <- false
			}
			body, _ := ioutil.ReadAll(res.Body)
			msg = string(body)
			utils.Log(fmt.Sprintf("Response: %s", msg))
			wg.Done()
		}()
	}
	go func() {
		for {
			select {
			case <-errorChannel:
				utils.Log("An error has occured")
			}
		}
	}()
	wg.Wait()
	endTime := time.Now()

	fmt.Printf("Execution time: %d seconds", endTime.Unix()-startTime.Unix())
}

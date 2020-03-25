package main

import (
	"fmt"
	"github.com/nyugoh/gopher-world/rabbitmq/utils"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

func main() {
	con, _ := utils.RedisConnect()
	channel, _ := utils.GetChannel(con)
	channel.Qos(
		1,
		0,
		false,
	)
	queue, _ := utils.DeclareChannelQueue(channel, utils.QueueName, true)
	Consume(channel, &queue)
}

func Consume(channel *amqp.Channel, queue *amqp.Queue) {
	messages, _ := utils.Consume(channel, queue.Name)

	go func() {
		i := 1
		for msg := range messages {
			fmt.Printf("Received message %d:%s\n", i, msg.Body)
			// Sleep for random time 1-5 seconds, simulates heavy work
			workTime := time.Duration(rand.Intn(5) + 1)
			// Send ack
			fmt.Println("Working for", workTime.Seconds(), "Seconds")
			time.Sleep(workTime * time.Second)
			i++
			fmt.Println("Done processing...", i)
			msg.Ack(false)
		}
	}()

	// Wait for messages
	forever := make(<-chan bool)
	log.Println("Waiting for messages....")
	<-forever
	log.Println("Exiting app...")
}

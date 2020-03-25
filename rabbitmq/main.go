package main

import (
	"fmt"
	"github.com/nyugoh/gopher-world/rabbitmq/utils"
	"github.com/streadway/amqp"
	"log"
	"strings"
	"time"
)

var (
	QueueName = "QBT_INV_ITEMS"
)

func main() {

	con, _ := utils.RedisConnect()
	defer con.Close()

	channel, _ := utils.GetChannel(con)

	queue, _ := utils.DeclareChannelQueue(channel, QueueName)

	message := "Jane Doe"

	// Push messages to queue
	go func() {
		for i :=1;;i++ {
			message = fmt.Sprintf("%s-%d", strings.Split(message, "-")[0], time.Now().Unix())
			err := channel.Publish(
				"",
				queue.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(message),
				})
			if err != nil {
				msg := "Error occurred when publishing message"
				fmt.Println(msg, err.Error())
			}
			fmt.Printf("Message %d published successfully...\n", i)

			time.Sleep(time.Second * 1)
		}
	}()

	Consume()
}

func Consume() {
	con, err := utils.RedisConnect()

	if Fail(err, "unable to connect to MQ") {
		return
	}

	channel, err := con.Channel()
	if Fail(err, "unable to create channel") {
		return
	}

	// Declare queue
	queue, err := channel.QueueDeclare(QueueName, false, false, false, false, nil)
	if Fail(err, "unable to create a queue") {
		return
	}

	forever := make(<-chan bool)

	msgs, err := channel.Consume(queue.Name, "", false, false, false, false, nil)

	go func() {
		i :=1
		for msg := range msgs {
			fmt.Printf("Received message %d:%s\n", i, msg.Body)
			time.Sleep(time.Second * 2)
			i++
		}
	}()

	log.Println("Waiting for messages....")
	<-forever
	log.Println("Exiting app...")
}

func Fail(err error, message string) bool {
	if err != nil {
		log.Fatal(message, err.Error())
		return true
	}
	return false
}

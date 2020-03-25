package main

import (
	"fmt"
	"github.com/nyugoh/gopher-world/rabbitmq/utils"
	"github.com/streadway/amqp"
	"time"
)

func main() {
	con, _ := utils.RedisConnect()
	channel, _ := utils.GetChannel(con)
	queue, _ := utils.DeclareChannelQueue(channel, utils.QueueName, true)

	Publish(channel, &queue)

	forever := make(<- chan bool)
	<- forever
}

func Publish(channel *amqp.Channel,queue *amqp.Queue) {
	message := "Jane Doe"
	for i := 1; ; i++ {
		message = fmt.Sprintf("%s-%d", "Message", i)
		err := channel.Publish(
			"",
			queue.Name,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent, // Allow persistent messages across restarts
				ContentType: "text/plain",
				Body:        []byte(message),
			})
		if err != nil {
			msg := "Error occurred when publishing message"
			fmt.Println(msg, err.Error())
		}
		fmt.Printf("Message %d published successfully...\n", i)
		// Pushes jobs every 1 second
		time.Sleep(time.Microsecond * 500)
	}
}


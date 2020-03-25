package utils

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func init()  {
	fmt.Println("Init utils")
}

func RedisConnect() (connection *amqp.Connection, err error)  {
	connection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		msg := "Unable to connect to Rabbit MQ"
		log.Fatal(msg, err.Error())
		return nil, err
	}
	return
}

func GetChannel(connection *amqp.Connection) (channel *amqp.Channel, err error)  {
	channel, err = connection.Channel()
	if err != nil {
		msg := "Unable to create a channel"
		log.Fatal(msg, err.Error())
	}
	return
}

func DeclareChannelQueue(channel *amqp.Channel, queueName string) (queue amqp.Queue, err error) {
	queue, err = channel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		msg := "Unable to declare a queue"
		log.Fatal(msg, err.Error())
	}
	return
}

func Consume(channel *amqp.Channel, queueName string) (payload <- chan amqp.Delivery, err error) {
	payload, err = channel.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		msg := "Unable to declare a queue"
		log.Fatal(msg, err.Error())
	}
	return
}
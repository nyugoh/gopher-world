package gokit

import (
	"fmt"

	"github.com/streadway/amqp"
)

// MQConnect connects to rabbitmq and returns a pointer to a connection/ error
// this function should be called at the initialization of the app and the connection stored
// Channels should be created/closed on demand from this one connection.
func MQConnect(user, password, host, port, vhost string) (connection *amqp.Connection, err error) {
	connectionString := fmt.Sprintf("amqp://%v:%v@%v:%v%v", user, password, host, port, vhost)
	Log("%s Connecting to rabbitMq...", LogPrefix)
	connection, err = amqp.Dial(connectionString)
	if err != nil {
		err = fmt.Errorf("%s failed to connect to RabbitMQ, err: %v", LogPrefix, err)
		return nil, err
	}
	Log("%s Connected to rabbitMq successfully", LogPrefix)
	return
}

package queue

import (
	"log"

	"github.com/streadway/amqp"
)

//AMQConfig basic struct
type AMQConfig struct {
	AMQPConnectionURL string
	Connection        *amqp.Connection
}

func newAMQConfig(connectionString string) (AMQConfig, error) {
	var amqconfig AMQConfig
	var err error
	amqconfig.AMQPConnectionURL = connectionString

	return amqconfig, err
}

//ConnectQueue will connect to the URL
func (A *AMQConfig) ConnectQueue() {
	var err error
	A.Connection, err = amqp.Dial(A.AMQPConnectionURL)
	if err != nil {
		log.Println(err)
	}
}

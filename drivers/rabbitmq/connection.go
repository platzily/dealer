package rabbitmq

import (
	"github.com/platzily/dealer/config"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func NewConnection() (conn *amqp.Connection, err error) {

	env := config.ReadRabbitMQConfig()
	conn, err = amqp.Dial(env.URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	return conn, nil
}

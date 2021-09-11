package rabbitmq

import (
	"github.com/platzily/dealer/config/rabbitmq"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func NewConnection() (conn *amqp.Connection, err error) {

	url := rabbitmq.ReadRabbitMQConfig()
	conn, err = amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	return conn, nil
}

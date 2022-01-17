package rabbitmq

import (
	"github.com/platzily/dealer/config"
	"github.com/platzily/dealer/domains"
	"github.com/streadway/amqp"
)

const MAIN_EVENTS_QUEUE = "main_queue"

var env *config.RabbitMQConfig = config.ReadRabbitMQConfig()

type EventQueue struct {
	conn *amqp.Connection
}

func New(rabbitConnection *amqp.Connection) domains.EventQueue {
	return &EventQueue{
		conn: rabbitConnection,
	}
}

func (eq *EventQueue) Subscribe(eventType string, callback func()) error {

	return nil
}

func (eq *EventQueue) Publish(event domains.Event) error {

	return nil
}

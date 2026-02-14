package rabbit

import (
	"github.com/buemura/event-driven-commerce/svc-order/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var QueueDeclareList []string = []string{
	"order.completed",
	"order.completed.dlq",
}

var QueueConsumerList []string = []string{
	"order.completed",
}

func DeclareQueueList() {
	conn, err := amqp.Dial(config.BROKER_URL)
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	for _, q := range QueueDeclareList {
		_, err := ch.QueueDeclare(
			q,     // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			log.Panicf("Failed to declare queue %s: %s", q, err)
		}
	}
}

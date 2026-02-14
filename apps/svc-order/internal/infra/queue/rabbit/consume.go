package rabbit

import (
	"log"

	"github.com/buemura/event-driven-commerce/svc-order/config"
	"github.com/buemura/event-driven-commerce/svc-order/internal/infra/queue/controller"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumeIn struct {
	Queue string
}

func Consume(in *ConsumeIn) {
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

	msgs, err := ch.Consume(
		in.Queue, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		log.Panicf("Failed to register a consumer: %s", err)
	}

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("\n")
			handleMessage(d)
		}
	}()

	log.Printf("RabbitMQ Consumer running for: Queue = %s", in.Queue)
	<-forever
}

func handleMessage(d amqp.Delivery) {
	switch d.RoutingKey {
	case "order.completed":
		controller.UpdateOrderStatus(string(d.Body))
	}
}

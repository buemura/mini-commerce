package queue

import (
	"context"
	"log"
	"time"

	"github.com/buemura/event-driven-commerce/svc-order/config"
	"github.com/buemura/event-driven-commerce/svc-order/internal/application/contracts"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PublishIn struct {
	Exchange   string
	RoutingKey string
	Payload    string
}

func Publish(in *PublishIn) {
	p := NewRabbitPublisher()
	err := p.Publish(&contracts.PublishInput{
		Exchange:   in.Exchange,
		RoutingKey: in.RoutingKey,
		Payload:    in.Payload,
	})
	if err != nil {
		log.Printf("[Queue][Publish] - Error: %s", err)
	}
}

type RabbitPublisher struct{}

func NewRabbitPublisher() *RabbitPublisher {
	return &RabbitPublisher{}
}

func (p *RabbitPublisher) Publish(in *contracts.PublishInput) error {
	conn, err := amqp.Dial(config.BROKER_URL)
	if err != nil {
		log.Printf("[Queue][Publish] - Failed to connect to RabbitMQ: %s", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("[Queue][Publish] - Failed to open a channel: %s", err)
		return err
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		in.Exchange,   // exchange
		in.RoutingKey, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(in.Payload),
		})
	if err != nil {
		log.Printf("[Queue][Publish] - Failed to publish a message: %s", err)
		return err
	}

	log.Printf("[Queue][Publish] - Sent message to %s\n", in.RoutingKey)
	return nil
}

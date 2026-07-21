package queue

import (
	"context"
	"log"
	"time"

	"github.com/buemura/event-driven-commerce/packages/tracing"
	"github.com/buemura/event-driven-commerce/mc-payment-service/config"
	"github.com/buemura/event-driven-commerce/mc-payment-service/internal/infra/util"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type PublishIn struct {
	Exchange    string
	Queue       string
	RountingKey string
	Payload     string
}

func Publish(ctx context.Context, in *PublishIn) {
	tracer := otel.Tracer("mc-payment-service")
	ctx, span := tracer.Start(ctx, "rabbitmq.publish "+in.RountingKey)
	defer span.End()

	conn, err := amqp.Dial(config.BROKER_URL)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	pubCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	headers := tracing.InjectAMQPHeaders(ctx)

	err = ch.PublishWithContext(pubCtx,
		in.Exchange,    // exchange
		in.RountingKey, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Headers:     headers,
			Body:        []byte(in.Payload),
		})
	util.FailOnError(err, "Failed to publish a message")
	log.Printf("[Queue][Publish] - Sent message to %s: \n", in.RountingKey)
}

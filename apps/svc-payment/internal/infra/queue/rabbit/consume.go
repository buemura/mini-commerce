package rabbit

import (
	"context"
	"log"

	"github.com/buemura/event-driven-commerce/packages/tracing"
	"github.com/buemura/event-driven-commerce/svc-payment/config"
	"github.com/buemura/event-driven-commerce/svc-payment/internal/infra/queue/controller"
	"github.com/buemura/event-driven-commerce/svc-payment/internal/infra/util"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type ConsumeIn struct {
	Queue string
}

func Consume(in *ConsumeIn) {
	conn, err := amqp.Dial(config.BROKER_URL)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
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
	util.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("\n")
			ctx := tracing.ExtractAMQPContext(context.Background(), d.Headers)
			tracer := otel.Tracer("svc-payment")
			ctx, span := tracer.Start(ctx, "rabbitmq.consume "+d.RoutingKey)
			handleMessage(ctx, d)
			span.End()
		}
	}()

	log.Printf("RabbitMQ Consumer running for: Queue = %s", in.Queue)
	<-forever
}

func handleMessage(ctx context.Context, d amqp.Delivery) {
	switch d.RoutingKey {
	case "order.create":
		controller.CreateOrder(ctx, string(d.Body))
	case "order.update":
		controller.UpdateOrder(ctx, string(d.Body))
	case "payment.create":
		controller.CreatePayment(ctx, string(d.Body))
	case "payment.process":
		controller.ProcessPayment(ctx, string(d.Body))
	}
}

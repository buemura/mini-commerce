package controller

import (
	"context"
	"encoding/json"
	"log"

	"github.com/buemura/event-driven-commerce/svc-payment/internal/application/usecase"
	"github.com/buemura/event-driven-commerce/svc-payment/internal/domain/order"
	"github.com/buemura/event-driven-commerce/svc-payment/internal/domain/payment"
	"github.com/buemura/event-driven-commerce/svc-payment/internal/infra/database"
	"github.com/buemura/event-driven-commerce/svc-payment/internal/infra/queue"
)

func CreatePayment(ctx context.Context, payload string) {
	var in *payment.CreatePaymentIn
	err := json.Unmarshal([]byte(payload), &in)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("[QueueController][CreatePayment] - Init payment create for order:", in.OrderId)

	repo := database.NewPgxPaymentRepository()
	uc := usecase.NewPaymentCreateUsecase(repo)

	p, err := uc.Execute(ctx, in)
	if err != nil {
		log.Println("[QueueController][CreatePayment] - Error:", err.Error())
		queue.Publish(ctx, &queue.PublishIn{
			RountingKey: "payment.create.dlq",
			Payload:     payload,
		})
		return
	}
	log.Println("[QueueController][CreatePayment] - Successfully inserted payment for order:", in.OrderId)

	processPaymentPayload, _ := json.Marshal(&payment.CreatePaymentOut{
		OrderId: p.OrderId,
	})
	queue.Publish(ctx, &queue.PublishIn{
		RountingKey: "payment.process",
		Payload:     string(processPaymentPayload),
	})
}

func ProcessPayment(ctx context.Context, payload string) {
	var in *payment.ProcessPaymentIn
	err := json.Unmarshal([]byte(payload), &in)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("[QueueController][ProcessPayment] - Init payment processing for order:", in.OrderId)

	repo := database.NewPgxPaymentRepository()
	uc := usecase.NewPaymentProcessUsecase(repo)

	p, err := uc.Execute(ctx, in)
	if err != nil {
		log.Println("[QueueController][ProcessPayment] - Error:", err.Error())
		queue.Publish(ctx, &queue.PublishIn{
			RountingKey: "payment.create.dlq",
			Payload:     payload,
		})
		return
	}
	log.Println("[QueueController][ProcessPayment] - Successfully processed payment for order:", in.OrderId)

	if p.Status == payment.PaymentFailed {
		paymentCreate, _ := json.Marshal(&order.CreateOrderOut{
			OrderID: p.OrderId,
		})
		queue.Publish(ctx, &queue.PublishIn{
			RountingKey: "payment.create",
			Payload:     string(paymentCreate),
		})
		return
	}

	orderUpdatePayload, _ := json.Marshal(&order.UpdateOrderIn{
		OrderId: p.OrderId,
		Status:  order.StatusCompleted,
	})
	queue.Publish(ctx, &queue.PublishIn{
		RountingKey: "order.update",
		Payload:     string(orderUpdatePayload),
	})
	queue.Publish(ctx, &queue.PublishIn{
		RountingKey: "order.completed",
		Payload:     string(orderUpdatePayload),
	})
}

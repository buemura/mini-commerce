package controller

import (
	"context"
	"encoding/json"
	"log"

	"github.com/buemura/event-driven-commerce/mc-payment-service/internal/application/usecase"
	"github.com/buemura/event-driven-commerce/mc-payment-service/internal/domain/order"
	"github.com/buemura/event-driven-commerce/mc-payment-service/internal/infra/database"
	"github.com/buemura/event-driven-commerce/mc-payment-service/internal/infra/queue"
)

func CreateOrder(ctx context.Context, payload string) {
	var in *order.CreateOrderIn
	err := json.Unmarshal([]byte(payload), &in)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("[QueueController][CreateOrder] - Init order creation for order:", in.OrderId)

	repo := database.NewPgxOrderRepository()
	uc := usecase.NewOrderCreateUsecase(repo)
	o, err := uc.Execute(ctx, in)
	if err != nil {
		log.Println("[QueueController][CreateOrder] - Error:", err.Error())
		queue.Publish(ctx, &queue.PublishIn{
			RountingKey: "order.create.dlq",
			Payload:     payload,
		})
		return
	}
	log.Println("[QueueController][CreateOrder] - Successfully created order:", in.OrderId)

	paymentCreate, _ := json.Marshal(&order.CreateOrderOut{
		OrderID:       o.ID,
		Amount:        o.Amount,
		PaymentMethod: o.PaymentMethod,
	})
	queue.Publish(ctx, &queue.PublishIn{
		RountingKey: "payment.create",
		Payload:     string(paymentCreate),
	})
}

func UpdateOrder(ctx context.Context, payload string) {
	var in *order.UpdateOrderIn
	err := json.Unmarshal([]byte(payload), &in)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("[QueueController][UpdateOrder] - Init order update for order:", in.OrderId)

	repo := database.NewPgxOrderRepository()
	uc := usecase.NewOrderUpdateUsecase(repo)
	o, err := uc.Execute(ctx, in)
	if err != nil {
		log.Println("[QueueController][UpdateOrder] - Error:", err.Error())
		queue.Publish(ctx, &queue.PublishIn{
			RountingKey: "order.create.dlq",
			Payload:     payload,
		})
		return
	}
	log.Println(o)
	log.Println("[QueueController][UpdateOrder] - Successfully updated order:", in.OrderId)
}

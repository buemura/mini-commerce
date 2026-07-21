package controller

import (
	"context"
	"encoding/json"
	"log"

	"github.com/buemura/event-driven-commerce/mc-order-service/internal/domain/order"
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/infra/factory"
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/infra/queue"
)

func UpdateOrderStatus(ctx context.Context, payload string) {
	var in *order.UpdateOrderStatusIn
	err := json.Unmarshal([]byte(payload), &in)
	if err != nil {
		log.Printf("[QueueController][UpdateOrderStatus] - Failed to unmarshal payload: %s", err)
		return
	}
	log.Println("[QueueController][UpdateOrderStatus] - Init order status update for order:", in.OrderId)

	uc := factory.MakeUpdateOrderStatusUsecase()
	err = uc.Execute(ctx, in)
	if err != nil {
		log.Println("[QueueController][UpdateOrderStatus] - Error:", err.Error())
		queue.Publish(ctx, &queue.PublishIn{
			RoutingKey: "order.completed.dlq",
			Payload:    payload,
		})
		return
	}
	log.Println("[QueueController][UpdateOrderStatus] - Successfully updated order status for order:", in.OrderId)
}

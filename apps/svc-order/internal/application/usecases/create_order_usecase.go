package usecases

import (
	"context"
	"encoding/json"
	"log"

	"github.com/buemura/event-driven-commerce/svc-order/internal/application/contracts"
	"github.com/buemura/event-driven-commerce/svc-order/internal/domain/order"
)

type CreateOrderUsecase struct {
	repo           order.OrderRepository
	productService contracts.ProductService
	publisher      contracts.QueuePublisher
}

func NewCreateOrderUsecase(
	repo order.OrderRepository,
	productService contracts.ProductService,
	publisher contracts.QueuePublisher,
) *CreateOrderUsecase {
	return &CreateOrderUsecase{
		repo:           repo,
		productService: productService,
		publisher:      publisher,
	}
}

func (s *CreateOrderUsecase) Execute(ctx context.Context, in *order.CreateOrderIn) (*order.Order, error) {
	o, err := order.NewOrder(in)
	if err != nil {
		return nil, err
	}

	for _, p := range in.ProductList {
		_, err := s.productService.UpdateProductQuantity(ctx, p.ID, (p.Quantity * -1))
		if err != nil {
			return nil, err
		}
	}

	res, err := s.repo.Save(ctx, o)
	if err != nil {
		return nil, err
	}

	s.publishOrderCreated(ctx, res)

	return res, nil
}

func (s *CreateOrderUsecase) publishOrderCreated(ctx context.Context, o *order.Order) {
	payload := map[string]interface{}{
		"order_id":       o.ID,
		"amount":         o.TotalPrice,
		"payment_method": o.PaymentMethod,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[CreateOrderUsecase][publishOrderCreated] - Failed to marshal payload: %s", err)
		return
	}

	err = s.publisher.Publish(ctx, &contracts.PublishInput{
		RoutingKey: "order.create",
		Payload:    string(data),
	})
	if err != nil {
		log.Printf("[CreateOrderUsecase][publishOrderCreated] - Failed to publish order.create: %s", err)
	}
}

package database

import (
	"context"

	"github.com/buemura/event-driven-commerce/svc-payment/internal/domain/order"
)

type InMemoryOrderRepo struct {
	order []*order.Order
}

func NewInMemoryOrderRepo(order []*order.Order) *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		order: order,
	}
}

func (r *InMemoryOrderRepo) FindMany(_ context.Context, in *order.GetManyOrdersIn) (*order.OrderRepositoryPaginatedOut, error) {
	return &order.OrderRepositoryPaginatedOut{
		OrderList:  r.order,
		TotalCount: len(r.order),
	}, nil
}

func (r *InMemoryOrderRepo) FindById(_ context.Context, id string) (*order.Order, error) {
	var o *order.Order
	for _, v := range r.order {
		if v.ID == id {
			o = v
			break
		}
	}
	return o, nil
}

func (r *InMemoryOrderRepo) Save(_ context.Context, o *order.Order) (*order.Order, error) {
	r.order = append(r.order, o)
	return o, nil
}

func (r *InMemoryOrderRepo) Update(_ context.Context, id, status string) error {
	for _, v := range r.order {
		if v.ID == id {
			v.Status = order.OrderStatus(status)
			break
		}
	}
	return nil
}

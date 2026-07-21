package usecase

import (
	"context"

	"github.com/buemura/event-driven-commerce/mc-payment-service/internal/domain/order"
)

type OrderCreateUsecase struct {
	repo order.OrderRepository
}

func NewOrderCreateUsecase(repo order.OrderRepository) *OrderCreateUsecase {
	return &OrderCreateUsecase{
		repo: repo,
	}
}

func (u *OrderCreateUsecase) Execute(ctx context.Context, in *order.CreateOrderIn) (*order.Order, error) {
	o, err := order.NewOrder(in)
	if err != nil {
		return nil, err
	}

	_, err = u.repo.Save(ctx, o)
	if err != nil {
		return nil, err
	}

	return o, err
}

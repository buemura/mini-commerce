package usecase

import (
	"context"

	"github.com/buemura/event-driven-commerce/mc-payment-service/internal/domain/order"
)

type OrderUpdateUsecase struct {
	repo order.OrderRepository
}

func NewOrderUpdateUsecase(repo order.OrderRepository) *OrderUpdateUsecase {
	return &OrderUpdateUsecase{
		repo: repo,
	}
}

func (u *OrderUpdateUsecase) Execute(ctx context.Context, in *order.UpdateOrderIn) (*order.Order, error) {
	o, err := u.repo.FindById(ctx, in.OrderId)
	if err != nil {
		return nil, order.ErrOrderNotFound
	}

	o.Status = in.Status

	err = u.repo.Update(ctx, o.ID, string(o.Status))
	if err != nil {
		return nil, err
	}

	return o, nil
}

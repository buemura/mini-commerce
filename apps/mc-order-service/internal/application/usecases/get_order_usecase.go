package usecases

import (
	"context"

	"github.com/buemura/event-driven-commerce/mc-order-service/internal/domain/order"
)

type GetOrderUsecase struct {
	repo order.OrderRepository
}

func NewGetOrderUsecase(repo order.OrderRepository) *GetOrderUsecase {
	return &GetOrderUsecase{
		repo: repo,
	}
}

func (s *GetOrderUsecase) Execute(ctx context.Context, id string) (*order.Order, error) {
	ord, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if ord == nil {
		return nil, order.ErrOrderNotFound
	}
	return ord, nil
}

package usecases

import (
	"context"

	"github.com/buemura/event-driven-commerce/svc-order/internal/domain/order"
)

type UpdateOrderStatusUsecase struct {
	repo order.OrderRepository
}

func NewUpdateOrderStatusUsecase(repo order.OrderRepository) *UpdateOrderStatusUsecase {
	return &UpdateOrderStatusUsecase{
		repo: repo,
	}
}

func (s *UpdateOrderStatusUsecase) Execute(ctx context.Context, in *order.UpdateOrderStatusIn) error {
	o, err := s.repo.FindById(ctx, in.OrderId)
	if err != nil {
		return err
	}
	if o == nil {
		return order.ErrOrderNotFound
	}

	return s.repo.UpdateStatus(ctx, in.OrderId, in.Status)
}

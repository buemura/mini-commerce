package usecases

import (
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

func (s *UpdateOrderStatusUsecase) Execute(in *order.UpdateOrderStatusIn) error {
	o, err := s.repo.FindById(in.OrderId)
	if err != nil {
		return err
	}
	if o == nil {
		return order.ErrOrderNotFound
	}

	return s.repo.UpdateStatus(in.OrderId, in.Status)
}

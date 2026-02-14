package usecase

import (
	"github.com/buemura/event-driven-commerce/svc-payment/internal/domain/order"
)

type OrderCreateUsecase struct {
	repo order.OrderRepository
}

func NewOrderCreateUsecase(repo order.OrderRepository) *OrderCreateUsecase {
	return &OrderCreateUsecase{
		repo: repo,
	}
}

func (u *OrderCreateUsecase) Execute(in *order.CreateOrderIn) (*order.Order, error) {
	o, err := order.NewOrder(in)
	if err != nil {
		return nil, err
	}

	_, err = u.repo.Save(o)
	if err != nil {
		return nil, err
	}

	return o, err
}

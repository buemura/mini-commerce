package usecases

import (
	"context"

	"github.com/buemura/event-driven-commerce/mc-customer-service/internal/domain/customer"
)

type CustomerGetService struct {
	repo customer.CustomerRepository
}

func NewCustomerGetService(repo customer.CustomerRepository) *CustomerGetService {
	return &CustomerGetService{
		repo: repo,
	}
}

func (s *CustomerGetService) Execute(ctx context.Context, customerID string) (*customer.CustomerOut, error) {
	cust, err := s.repo.FindById(ctx, customerID)
	if err != nil {
		return nil, err
	}
	if cust == nil {
		return nil, customer.ErrCustomerNotFound
	}

	return &customer.CustomerOut{
		ID:    cust.ID,
		Name:  cust.Name,
		Email: cust.Email,
	}, nil
}

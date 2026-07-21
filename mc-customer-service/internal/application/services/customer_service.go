package services

import (
	"context"

	"github.com/buemura/event-driven-commerce/mc-customer-service/internal/application/contracts"
	"github.com/buemura/event-driven-commerce/mc-customer-service/internal/domain/customer"
)

type CustomerService struct {
	repo           customer.CustomerRepository
	passwordHasher contracts.PasswordHasher
	tokenGenerator contracts.TokenGenerator
}

func NewCustomerService(
	repo customer.CustomerRepository,
	passwordHasher contracts.PasswordHasher,
	tokenGenerator contracts.TokenGenerator,
) *CustomerService {
	return &CustomerService{
		repo:           repo,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
	}
}

func (s *CustomerService) Get(ctx context.Context, customerID string) (*customer.CustomerOut, error) {
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

func (s *CustomerService) Signin(ctx context.Context, in *customer.SigninCustomerIn) (*customer.SigninCustomerOut, error) {
	cust, err := s.repo.FindByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if cust == nil {
		return nil, customer.ErrCustomerInvalidCredential
	}

	isPassValid := s.passwordHasher.Compare(in.Password, cust.Password)
	if !isPassValid {
		return nil, customer.ErrCustomerInvalidCredential
	}

	accessToken, err := s.tokenGenerator.Generate(cust.ID)
	if err != nil {
		return nil, err
	}

	return &customer.SigninCustomerOut{
		AccessToken: accessToken,
		Customer: customer.CustomerOut{
			ID:    cust.ID,
			Name:  cust.Name,
			Email: cust.Email,
		},
	}, nil
}

func (s *CustomerService) Signup(ctx context.Context, in *customer.CreateCustomerIn) error {
	custExists, err := s.repo.FindByEmail(ctx, in.Email)
	if err != nil {
		return err
	}
	if custExists != nil {
		return customer.ErrCustomerAlreadyExists
	}

	hashed, err := s.passwordHasher.Hash(in.Password)
	if err != nil {
		return err
	}

	in.Password = hashed

	cust, err := customer.NewCustomer(in)
	if err != nil {
		return err
	}

	_, err = s.repo.Save(ctx, cust)
	if err != nil {
		return err
	}
	return nil
}

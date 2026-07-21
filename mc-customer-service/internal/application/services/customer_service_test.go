package services

import (
	"context"
	"errors"
	"testing"

	"github.com/buemura/event-driven-commerce/mc-customer-service/internal/domain/customer"
	"github.com/buemura/event-driven-commerce/mc-customer-service/internal/infra/adapters"
	"github.com/buemura/event-driven-commerce/mc-customer-service/internal/infra/database"

	"github.com/stretchr/testify/assert"
)

func makeService(customers []customer.Customer) *CustomerService {
	repo := database.NewInMemoryCustomerRepo(customers)
	hasher := adapters.NewStubPasswordHasher()
	tkGen := adapters.NewStubTokenGenerator()
	return NewCustomerService(repo, hasher, tkGen)
}

func TestCustomerServiceGet(t *testing.T) {
	service := makeService([]customer.Customer{
		{
			ID:       "existing_id",
			Name:     "existing_name",
			Email:    "existing_email",
			Password: "hashed:existing_password",
		},
	})

	t.Run("Customer not found", func(t *testing.T) {
		_, err := service.Get(context.Background(), "not_found_id")
		assert.Equal(t, err, customer.ErrCustomerNotFound)
		assert.NotNil(t, err)
	})

	t.Run("Customer get success", func(t *testing.T) {
		res, _ := service.Get(context.Background(), "existing_id")
		assert.Equal(t, &customer.CustomerOut{
			ID:    "existing_id",
			Name:  "existing_name",
			Email: "existing_email",
		}, res)
	})
}

func TestCustomerServiceSignin(t *testing.T) {
	service := makeService([]customer.Customer{
		{
			ID:       "existing_id",
			Name:     "existing_name",
			Email:    "existing_email",
			Password: "hashed:existing_password",
		},
	})

	t.Run("customer not found", func(t *testing.T) {
		input := &customer.SigninCustomerIn{
			Email:    "no_email",
			Password: "no_password",
		}

		_, err := service.Signin(context.Background(), input)
		assert.Equal(t, err, customer.ErrCustomerInvalidCredential)
		assert.NotNil(t, err)
	})

	t.Run("customer invalid password", func(t *testing.T) {
		input := &customer.SigninCustomerIn{
			Email:    "existing_email",
			Password: "wrong_password",
		}

		_, err := service.Signin(context.Background(), input)
		assert.Equal(t, err, customer.ErrCustomerInvalidCredential)
		assert.NotNil(t, err)
	})

	t.Run("customer signin success", func(t *testing.T) {
		input := &customer.SigninCustomerIn{
			Email:    "existing_email",
			Password: "existing_password",
		}

		res, err := service.Signin(context.Background(), input)
		assert.Equal(t, err, nil)
		assert.Equal(t, &customer.SigninCustomerOut{
			AccessToken: "existing_id",
			Customer: customer.CustomerOut{
				ID:    "existing_id",
				Name:  "existing_name",
				Email: "existing_email",
			},
		}, res)
	})
}

func TestCustomerServiceSignup(t *testing.T) {
	service := makeService([]customer.Customer{
		{
			ID:       "existing_id",
			Name:     "existing_name",
			Email:    "existing_email",
			Password: "existing_password",
		},
	})

	t.Run("customer already exists", func(t *testing.T) {
		input := &customer.CreateCustomerIn{
			Name:     "existing_name",
			Email:    "existing_email",
			Password: "existing_password",
		}

		err := service.Signup(context.Background(), input)
		assert.Equal(t, err, customer.ErrCustomerAlreadyExists)
		assert.NotNil(t, err)
	})

	t.Run("customer validation error", func(t *testing.T) {
		input := &customer.CreateCustomerIn{
			Name:     "",
			Email:    "any_email",
			Password: "any_password",
		}

		err := service.Signup(context.Background(), input)
		assert.Equal(t, err, errors.New("invalid customer name"))
		assert.NotNil(t, err)
	})
}

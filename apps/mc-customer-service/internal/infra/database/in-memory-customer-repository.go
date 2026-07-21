package database

import (
	"context"

	"github.com/buemura/event-driven-commerce/mc-customer-service/internal/domain/customer"
)

type InMemoryCustomerRepo struct {
	cus []customer.Customer
}

func NewInMemoryCustomerRepo(cus []customer.Customer) *InMemoryCustomerRepo {
	return &InMemoryCustomerRepo{
		cus: cus,
	}
}

func (r *InMemoryCustomerRepo) FindById(_ context.Context, id string) (*customer.Customer, error) {
	var c *customer.Customer
	for _, v := range r.cus {
		if v.ID == id {
			c = &v
			break
		}
	}
	return c, nil
}

func (r *InMemoryCustomerRepo) FindByEmail(_ context.Context, email string) (*customer.Customer, error) {
	var c *customer.Customer
	for _, v := range r.cus {
		if v.Email == email {
			c = &v
			break
		}
	}
	return c, nil
}

func (r *InMemoryCustomerRepo) Save(_ context.Context, cust *customer.Customer) (*customer.Customer, error) {
	r.cus = append(r.cus, *cust)
	return cust, nil
}

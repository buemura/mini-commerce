package customer

import "context"

type CustomerRepository interface {
	FindById(ctx context.Context, id string) (*Customer, error)
	FindByEmail(ctx context.Context, email string) (*Customer, error)
	Save(ctx context.Context, customer *Customer) (*Customer, error)
}

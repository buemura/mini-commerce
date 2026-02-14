package order

import "context"

type OrderRepositoryPaginatedOut struct {
	OrderList  []*Order
	TotalCount int
}

type OrderRepository interface {
	FindMany(ctx context.Context, in *GetManyOrdersIn) (*OrderRepositoryPaginatedOut, error)
	FindById(ctx context.Context, id string) (*Order, error)
	Save(ctx context.Context, o *Order) (*Order, error)
	UpdateStatus(ctx context.Context, id string, status string) error
}

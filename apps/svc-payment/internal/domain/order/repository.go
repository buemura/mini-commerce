package order

import "context"

type OrderRepositoryPaginatedOut struct {
	OrderList  []*Order
	TotalCount int
}

type OrderRepository interface {
	FindById(ctx context.Context, id string) (*Order, error)
	Save(ctx context.Context, o *Order) (*Order, error)
	Update(ctx context.Context, id, status string) error
}

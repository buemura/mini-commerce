package payment

import "context"

type PaymentRepository interface {
	FindById(ctx context.Context, id string) (*Payment, error)
	FindByOrderId(ctx context.Context, id string) ([]*Payment, error)
	FindPendingByOrderId(ctx context.Context, id string) (*Payment, error)
	Save(ctx context.Context, p *Payment) (*Payment, error)
	Update(ctx context.Context, id, status string) error
}

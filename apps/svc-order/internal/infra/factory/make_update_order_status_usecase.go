package factory

import (
	"github.com/buemura/event-driven-commerce/svc-order/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/svc-order/internal/infra/database"
)

func MakeUpdateOrderStatusUsecase() *usecases.UpdateOrderStatusUsecase {
	repo := database.NewPgxOrderRepository(database.Conn)
	usecase := usecases.NewUpdateOrderStatusUsecase(repo)
	return usecase
}

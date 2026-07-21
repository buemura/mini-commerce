package factory

import (
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/infra/database"
)

func MakeGetOrderUsecase() *usecases.GetOrderUsecase {
	repo := database.NewPgxOrderRepository(database.Conn)
	usecase := usecases.NewGetOrderUsecase(repo)
	return usecase
}

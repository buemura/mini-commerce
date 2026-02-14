package factory

import (
	"github.com/buemura/event-driven-commerce/svc-order/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/svc-order/internal/infra/database"
)

func MakeGetManyOrdersUsecase() *usecases.GetManyOrdersUsecase {
	repo := database.NewPgxOrderRepository(database.Conn)
	usecase := usecases.NewGetManyOrdersUsecase(repo)
	return usecase
}

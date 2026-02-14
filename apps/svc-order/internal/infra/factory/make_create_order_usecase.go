package factory

import (
	"github.com/buemura/event-driven-commerce/svc-order/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/svc-order/internal/infra/database"
	"github.com/buemura/event-driven-commerce/svc-order/internal/infra/grpc/client"
	"github.com/buemura/event-driven-commerce/svc-order/internal/infra/queue"
)

func MakeCreateOrderUsecase() *usecases.CreateOrderUsecase {
	repo := database.NewPgxOrderRepository(database.Conn)
	pService := client.NewProductServiceClient()
	publisher := queue.NewRabbitPublisher()
	usecase := usecases.NewCreateOrderUsecase(repo, pService, publisher)
	return usecase
}

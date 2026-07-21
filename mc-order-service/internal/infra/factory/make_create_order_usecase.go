package factory

import (
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/infra/database"
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/infra/grpc/client"
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/infra/queue"
)

func MakeCreateOrderUsecase() *usecases.CreateOrderUsecase {
	repo := database.NewPgxOrderRepository(database.Conn)
	pService := client.NewProductServiceClient()
	publisher := queue.NewRabbitPublisher()
	usecase := usecases.NewCreateOrderUsecase(repo, pService, publisher)
	return usecase
}

package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/buemura/event-driven-commerce/packages/metrics"
	"github.com/buemura/event-driven-commerce/packages/pb"
	"github.com/buemura/event-driven-commerce/packages/tracing"
	"github.com/buemura/event-driven-commerce/mc-order-service/config"
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/infra/database"
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/infra/grpc/server/controllers"
	"github.com/buemura/event-driven-commerce/mc-order-service/internal/infra/queue/rabbit"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func init() {
	config.LoadEnv()
	database.Connect()
	rabbit.DeclareQueueList()
}

func main() {
	ctx := context.Background()
	tp, err := tracing.InitTracer(ctx, "mc-order-service")
	if err != nil {
		log.Fatalf("Failed to initialize tracer: %v", err)
	}
	defer tp.Shutdown(ctx)

	mp, err := metrics.InitMeter(ctx, "mc-order-service")
	if err != nil {
		log.Fatalf("Failed to initialize meter: %v", err)
	}
	defer mp.Shutdown(ctx)

	metricsServer := metrics.Serve()
	defer metricsServer.Shutdown(ctx)

	port := ":" + config.GRPC_PORT
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot create grpc listener: %s", err)
	}

	s := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	pb.RegisterOrderServiceServer(s, &controllers.OrderController{})

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to server grpc: %s", err)
		}
	}()

	for _, q := range rabbit.QueueConsumerList {
		go rabbit.Consume(&rabbit.ConsumeIn{
			Queue: q,
		})
	}

	log.Println("gRPC Server running at", port, "...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop

	log.Println("Stopping gRPC Server...")
	s.GracefulStop()
	log.Println("gRPC Server stopped")
}

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
	"github.com/buemura/event-driven-commerce/mc-product-service/config"
	"github.com/buemura/event-driven-commerce/mc-product-service/internal/infra/database"
	"github.com/buemura/event-driven-commerce/mc-product-service/internal/infra/grpc/controllers"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func init() {
	config.LoadEnv()
	database.Connect()
}

func main() {
	ctx := context.Background()
	tp, err := tracing.InitTracer(ctx, "mc-product-service")
	if err != nil {
		log.Fatalf("Failed to initialize tracer: %v", err)
	}
	defer tp.Shutdown(ctx)

	mp, err := metrics.InitMeter(ctx, "mc-product-service")
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
	pb.RegisterProductServiceServer(s, &controllers.ProductController{})

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to server grpc: %s", err)
		}
	}()

	log.Println("gRPC Server running at", port, "...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop

	log.Println("Stopping gRPC Server...")
	s.GracefulStop()
	log.Println("gRPC Server stopped")
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buemura/event-driven-commerce/mc-api-gateway/config"
	"github.com/buemura/event-driven-commerce/mc-api-gateway/internal/infra/http/router"
	"github.com/buemura/event-driven-commerce/packages/tracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func init() {
	config.LoadEnv()
}

func main() {
	ctx := context.Background()
	tp, err := tracing.InitTracer(ctx, "mc-api-gateway")
	if err != nil {
		log.Fatalf("Failed to initialize tracer: %v", err)
	}
	defer tp.Shutdown(ctx)

	server := echo.New()
	server.Use(middleware.CORS())
	server.Use(otelecho.Middleware("mc-api-gateway"))

	router.SetupRouters(server)

	port := ":" + config.PORT

	go func() {
		if err := server.Start(port); err != nil && http.ErrServerClosed != err {
			panic(err)
		}
	}()

	log.Println("HTTP Server running at", port, "...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("Stopping HTTP Server...")

	if err := server.Shutdown(shutdownCtx); err != nil {
		panic(err)
	}
	log.Println("HTTP Server stopped")
}

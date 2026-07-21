package main

import (
	"context"
	"log"
	"sync"

	"github.com/buemura/event-driven-commerce/packages/metrics"
	"github.com/buemura/event-driven-commerce/packages/tracing"
	"github.com/buemura/event-driven-commerce/mc-payment-service/config"
	"github.com/buemura/event-driven-commerce/mc-payment-service/internal/infra/database"
	"github.com/buemura/event-driven-commerce/mc-payment-service/internal/infra/queue/rabbit"
)

func init() {
	config.LoadEnv()
	database.Connect()
	rabbit.DeclareQueueList()
}

func main() {
	ctx := context.Background()
	tp, err := tracing.InitTracer(ctx, "mc-payment-service")
	if err != nil {
		log.Fatalf("Failed to initialize tracer: %v", err)
	}
	defer tp.Shutdown(ctx)

	mp, err := metrics.InitMeter(ctx, "mc-payment-service")
	if err != nil {
		log.Fatalf("Failed to initialize meter: %v", err)
	}
	defer mp.Shutdown(ctx)

	metricsServer := metrics.Serve()
	defer metricsServer.Shutdown(ctx)

	var wg sync.WaitGroup
	wg.Add(len(rabbit.QueueConsumerList))

	for _, q := range rabbit.QueueConsumerList {
		go func() {
			defer wg.Done()
			rabbit.Consume(&rabbit.ConsumeIn{
				Queue: q,
			})
		}()
	}

	wg.Wait()
}

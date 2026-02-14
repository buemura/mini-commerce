package tracing

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type amqpCarrier amqp.Table

func (c amqpCarrier) Get(key string) string {
	if val, ok := (amqp.Table)(c)[key]; ok {
		if s, ok := val.(string); ok {
			return s
		}
	}
	return ""
}

func (c amqpCarrier) Set(key, value string) {
	(amqp.Table)(c)[key] = value
}

func (c amqpCarrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

func InjectAMQPHeaders(ctx context.Context) amqp.Table {
	headers := amqp.Table{}
	otel.GetTextMapPropagator().Inject(ctx, amqpCarrier(headers))
	return headers
}

func ExtractAMQPContext(ctx context.Context, headers amqp.Table) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, amqpCarrier(headers))
}

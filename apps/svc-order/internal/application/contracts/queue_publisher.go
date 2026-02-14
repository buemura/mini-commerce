package contracts

import "context"

type PublishInput struct {
	Exchange   string
	Queue      string
	RoutingKey string
	Payload    string
}

type QueuePublisher interface {
	Publish(ctx context.Context, in *PublishInput) error
}

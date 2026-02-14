package contracts

type PublishInput struct {
	Exchange   string
	Queue      string
	RoutingKey string
	Payload    string
}

type QueuePublisher interface {
	Publish(in *PublishInput) error
}

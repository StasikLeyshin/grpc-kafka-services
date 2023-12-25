package kafka

import "context"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type KafkaProducer interface {
	SendMessage(ctx context.Context, key []byte, value []byte) error
}

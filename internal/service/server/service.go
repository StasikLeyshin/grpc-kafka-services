package server

import (
	def "github.com/StasikLeyshin/grpc-kafka-services/internal/service"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/kafka"
)

var _ def.ServerService = (*service)(nil)

type service struct {
	producer kafka.KafkaProducer
}

func NewService(
	producer kafka.KafkaProducer,
) *service {
	return &service{
		producer: producer,
	}
}

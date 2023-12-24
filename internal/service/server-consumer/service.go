package server_consumer

import (
	"github.com/StasikLeyshin/grpc-kafka-services/internal/repository"
	def "github.com/StasikLeyshin/grpc-kafka-services/internal/service"
)

var _ def.ServerConsumerService = (*service)(nil)

type service struct {
	serverRepository repository.ServerRepository
}

func NewService(
	serverRepository repository.ServerRepository,
) *service {
	return &service{
		serverRepository: serverRepository,
	}
}

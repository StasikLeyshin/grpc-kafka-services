package server

import (
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service"
	desc "github.com/StasikLeyshin/grpc-kafka-services/pkg/server_v1"
)

type Implementation struct {
	desc.UnimplementedManagerServiceServer
	serverService service.ServerService
}

func NewImplementation(serverService service.ServerService) *Implementation {
	return &Implementation{
		serverService: serverService,
	}
}

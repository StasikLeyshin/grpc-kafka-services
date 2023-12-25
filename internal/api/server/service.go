package server

import (
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service"
	desc "github.com/StasikLeyshin/grpc-kafka-services/pkg/server_v1"
	"github.com/sirupsen/logrus"
)

type Implementation struct {
	desc.UnimplementedManagerServiceServer
	serverService service.ServerService
	logger        *logrus.Logger
}

func NewImplementation(serverService service.ServerService, logger *logrus.Logger) *Implementation {
	return &Implementation{
		serverService: serverService,
		logger:        logger,
	}
}

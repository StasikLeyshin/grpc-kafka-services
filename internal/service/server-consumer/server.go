package server_consumer

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/models"
)

func (s *service) Start() error {
	err := s.serverRepository.ServerAutoMigrate()
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateServer(ctx context.Context, server *models.Server, userUUID string) error {

	user := &models.User{
		UUID: userUUID,
	}

	err := s.serverRepository.CreateServer(server, user)
	if err != nil {
		return err
	}
	return nil
}

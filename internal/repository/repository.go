package repository

import "github.com/StasikLeyshin/grpc-kafka-services/internal/models"

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type ServerRepository interface {
	ServerAutoMigrate() error
	CreateServer(server *models.Server, user *models.User) error
	GetServer(serverUuid string) (*models.Server, error)
}

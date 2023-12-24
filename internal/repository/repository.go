package repository

import "github.com/StasikLeyshin/grpc-kafka-services/internal/models"

type ServerRepository interface {
	ServerAutoMigrate() error
	CreateServer(server *models.Server, user *models.User) error
}

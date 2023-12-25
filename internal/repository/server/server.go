package server

import (
	"fmt"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ServerAutoMigrate() error {
	err := r.db.AutoMigrate(&models.Server{}, &models.User{})
	if err != nil {
		return fmt.Errorf("error AutoMigrate: %v", err)
	}
	return nil
}

func (r *repository) CreateServer(server *models.Server, user *models.User) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&server).Error; err != nil {
			return err
		}

		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetServer(serverUuid string) (*models.Server, error) {
	var server models.Server

	err := r.db.First(&server, "uuid = ?", serverUuid)

	if err != nil {
		return nil, fmt.Errorf("server not found")
	}

	return &server, nil
}

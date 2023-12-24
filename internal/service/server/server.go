package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/models"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server/model"
	"github.com/google/uuid"
	"time"
)

func (s *service) CreateServer(ctx context.Context, serverRequest *model.CreateServerRequest) (*model.CreateServerResponse, error) {
	userUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generation user UUID: %v", err)
	}

	serverUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generation server UUID: %v", err)
	}

	server := models.Server{
		UUID:      serverUUID.String(),
		Host:      serverRequest.Host,
		Port:      serverRequest.Port,
		Name:      serverRequest.Name,
		CreatedAt: time.Now(),
	}

	key := []byte(userUUID.String())
	value, err := json.Marshal(server)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the request: %v", err)
	}

	err = s.producer.SendMessage(ctx, key, value)
	if err != nil {
		return nil, fmt.Errorf("failed to send message kafka: %v", err)
	}

	return &model.CreateServerResponse{
		Status: "ok",
	}, nil
}

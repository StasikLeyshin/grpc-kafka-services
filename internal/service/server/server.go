package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server/model"
	"github.com/google/uuid"
	"log"
)

func (s *service) CreateServer(ctx context.Context, serverRequest *model.CreateServerRequest) (*model.CreateServerResponse, error) {
	userUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generation server UUID: %v", err)
	}

	key := []byte(userUUID.String())
	value, err := json.Marshal(serverRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the request: %v", err)
	}

	err = s.producer.SendMessage(ctx, key, value)
	if err != nil {
		log.Printf("failed to send message kafka: %v\n", err)
		return nil, fmt.Errorf("failed to send message kafka: %v", err)
	}

	return &model.CreateServerResponse{
		Status: "ok",
	}, nil
}

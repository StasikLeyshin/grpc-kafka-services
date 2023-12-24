package converter

import (
	"encoding/json"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/models"
)

func ToServerFromKafka(msg []byte) *models.Server {
	var server models.Server

	_ = json.Unmarshal(msg, &server)

	return &server
}

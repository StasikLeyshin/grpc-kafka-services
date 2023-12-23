package startup

import (
	logger "github.com/StasikLeyshin/grpc-kafka-services/pkg"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	return logger.NewLogger()
}

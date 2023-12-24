package startup

import (
	"fmt"
	//"github.com/StasikLeyshin/grpc-kafka-services/internal/server/grpc"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	GrpcConfig  GrpcConfig     `yaml:"grpc"`
	KafkaConfig KafkaConfig    `yaml:"kafka"`
	Database    DatabaseConfig `yaml:"database"`
}

func NewConfig(configFile string) (*Config, error) {
	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("reading file error: %w", err)
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(rawYAML, cfg); err != nil {
		return nil, fmt.Errorf("yaml parsing error: %w", err)
	}

	return cfg, nil
}

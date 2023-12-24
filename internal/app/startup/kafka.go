package startup

import "github.com/segmentio/kafka-go"

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
}

func NewKafkaProducer(config KafkaConfig) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(config.Brokers...),
		Topic:    config.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func NewKafkaConsumer(config KafkaConfig) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   config.Brokers,
		Topic:     config.Topic,
		Partition: 0,
	})
}

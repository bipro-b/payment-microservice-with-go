package kafka

import (
	"context"
	"encoding/json"

	"github.com/bipro/payment-service/pkg/model"

	"github.com/segmentio/kafka-go"
)

var writer = kafka.NewWriter(kafka.WriterConfig{
	Brokers: []string{"kafka:9092"},
	Topic:   "payment-topic",
})

func Produce(p model.PaymentRequest) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return writer.WriteMessages(context.Background(), kafka.Message{Value: data})
}

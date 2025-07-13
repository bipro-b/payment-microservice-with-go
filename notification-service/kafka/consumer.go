package kafka

import (
	"context"
	"encoding/json"
	"log"

	"notification-service/grpc"
	pb "notification-service/proto"

	"github.com/segmentio/kafka-go"
)

// Consume starts consuming messages from the Kafka topic and handles them concurrently.
func Consume() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "payment-topic",
		GroupID: "notification-group",
	})

	defer reader.Close()
	log.Println("📡 Kafka consumer started...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("❌ Kafka read error: %v\n", err)
			continue
		}

		var p pb.PaymentRequest
		if err := json.Unmarshal(msg.Value, &p); err != nil {
			log.Printf("❌ JSON unmarshal error: %v\n", err)
			continue
		}

		log.Printf("✅ Received payment from user %s for %.2f via %s\n", p.UserId, p.Amount, p.Method)

		// Handle in a goroutine for concurrency
		go grpc.HandleNotification(&p)
	}
}

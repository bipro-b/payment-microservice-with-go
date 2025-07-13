package main

import (
	"log"

	"notification-service/kafka"
)

func main() {
	log.Println("🔔 Notification service started...")
	kafka.Consume()
}

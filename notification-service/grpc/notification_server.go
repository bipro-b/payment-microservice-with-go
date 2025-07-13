package grpc

import (
	"log"
	pb "notification-service/proto" // ✅ Import the protobuf types
)

func HandleNotification(p *pb.PaymentRequest) {
	log.Printf("📣 Notify user %s: Paid $%.2f via %s", p.UserId, p.Amount, p.Method)
}

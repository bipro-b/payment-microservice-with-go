

---

## 📦 Go Payment + Notification Microservices

A modern microservice architecture built with:

* 🔁 **Kafka** for async communication
* 💬 **gRPC + Protobuf** for contract-first messaging
* 🧠 **Goroutines** for concurrent processing
* 🧱 **Docker** + `docker-compose` for container orchestration
* ✨ Fully written in **Go**

---

## 🚀 Architecture Overview

```text
[Client/Frontend]
      |
      ↓  REST API (Gin)
┌──────────────────────┐
│   Payment Service     │
│ - Exposes /pay route  │
│ - Sends message to    │
│   Kafka "payment-topic"
└─────────▲────────────┘
          |
          ↓  Kafka (Async)
┌─────────────────────────────┐
│ Notification Service        │
│ - Consumes "payment-topic"  │
│ - Decodes PaymentRequest    │
│ - Passes it to              │
│   grpc.HandleNotification   │
└─────────────────────────────┘
```

---

## 📁 Folder Structure

```
payment-micro/
├── proto/                         # Protobuf contracts
│   └── payment.proto
│
├── payment-service/              # REST API service
│   ├── main.go
│   ├── kafka/producer.go
│   ├── handler/payment.go
│   └── Dockerfile
│
├── notification-service/         # Kafka + gRPC service
│   ├── main.go
│   ├── kafka/consumer.go
│   ├── grpc/notification_server.go
│   ├── proto/payment.pb.go       # Generated from .proto
│   ├── proto/payment_grpc.pb.go  # Generated from .proto
│   └── Dockerfile
│
├── docker-compose.yml
└── README.md
```

---

## 📦 Technologies Used

| Feature             | Tech Stack               |
| ------------------- | ------------------------ |
| Language            | Go 1.20+                 |
| API Framework       | Gin (in payment-service) |
| Messaging Queue     | Kafka + Zookeeper        |
| Async Communication | Kafka Topics             |
| Protocol Layer      | gRPC + Protobuf          |
| Concurrency         | Goroutines               |
| Containerization    | Docker, Docker Compose   |

---

## 🔌 API Endpoint (payment-service)

### `POST /pay`

Sends a payment request to Kafka.

#### ✅ Sample Request

```json
POST /pay
Content-Type: application/json

{
  "userId": "abc123",
  "amount": 59.99,
  "method": "credit"
}
```

---

## 🧪 Kafka Topic (payment-topic)

When a payment is received, a JSON message is pushed to:

```
Topic: payment-topic
```

Example Kafka message:

```json
{
  "userId": "abc123",
  "amount": 59.99,
  "method": "credit"
}
```

---

## 📬 gRPC Message Handling (notification-service)

The `notification-service` reads the Kafka message and unmarshals it into:

```go
*proto.PaymentRequest
```

Then it passes it to:

```go
grpc.HandleNotification(*pb.PaymentRequest)
```

---

## 📜 Protobuf Contract (proto/payment.proto)

```proto
syntax = "proto3";

package payment;

message PaymentRequest {
  string userId = 1;
  double amount = 2;
  string method = 3;
}

message Notification {
  string userId = 1;
  string message = 2;
}
```

---

## 🐳 Run With Docker Compose

```bash
docker-compose up --build
```

---

## 🧪 Test the API

```bash
curl -X POST http://localhost:8080/pay \
  -H "Content-Type: application/json" \
  -d '{"userId": "u001", "amount": 49.99, "method": "paypal"}'
```

---

## ✅ What We Implemented

✅ REST → Kafka producer
✅ Kafka consumer with goroutines
✅ Protobuf messages (`PaymentRequest`, `Notification`)
✅ gRPC-style function call internally
✅ Full module-based structure
✅ Dockerized environment
✅ Error handling & logging
✅ Code separation (`handler`, `kafka`, `grpc`)
✅ Multistage flow to demonstrate service-to-service messaging

---

## 🧠 Future Extensions (Optional)

* 🔁 Replace internal `HandleNotification()` with full gRPC server + client call
* 📬 Add response-acknowledgment via gRPC
* 💾 Persist payments/notifications in DB
* 🔐 Add Auth middleware (JWT or OAuth2)
* 📈 Add metrics with Prometheus/Grafana
* 📬 Integrate real-world notification channels (Email, SMS)

---

## 👨‍💻 Author

**Built by:** [@bipro](https://github.com/bipro)

---


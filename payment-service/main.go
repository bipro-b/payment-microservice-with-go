package main

import (
	"log"
	"net/http"

	"github.com/bipro/payment-service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/pay", handler.HandlePayment)
	log.Println("🚀 Payment service running on :8080")
	http.ListenAndServe(":8080", r)
}

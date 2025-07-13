package handler

import (
	"net/http"

	"github.com/bipro/payment-service/kafka"
	"github.com/bipro/payment-service/pkg/model"
	"github.com/gin-gonic/gin"
)

func HandlePayment(c *gin.Context) {
	var payment model.PaymentRequest
	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := kafka.Produce(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kafka failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment submitted"})
}

package model

type PaymentRequest struct {
	UserId string  `json:"userId"`
	Amount float64 `json:"amount"`
	Method string  `json:"method"`
}

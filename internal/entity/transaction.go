package entity

import "time"

type TransactionType int

const (
	Invoice  TransactionType = 1
	Withdraw TransactionType = 2
)

type TransactionStatus int

const (
	Error   TransactionStatus = 1
	Success TransactionStatus = 2
	Created TransactionStatus = 3
)

type Currency string

const (
	USDT Currency = "USDT"
	EUR  Currency = "EUR"
	RUB  Currency = "RUB"
)

func IsCurrencyValid(currency string) bool {
	return true
}

type CardBalance struct {
	CardNumber      int      `json:"cardNumber"`
	ActualBalance   float64  `json:"actualBalance"`
	PendingInvoice  float64  `json:"pendingInvoice"`
	PendingWithdraw float64  `json:"pendingWithdraw"`
	Currency        Currency `json:"currency"`
}

type Card struct {
	Number    int       `json:"number"`
	UserID    int       `json:"userId"`
	Currency  Currency  `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Transaction struct {
	ID         int               `json:"id"`
	UserID     int               `json:"userId"`
	CardNumber int               `json:"cardId"`
	Type       TransactionType   `json:"type"`
	Currency   Currency          `json:"currency"`
	Amount     float64           `json:"amount"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
	Status     TransactionStatus `json:"transactionStatus"`
}

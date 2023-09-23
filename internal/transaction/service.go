package transaction

import (
	"errors"
	"sync"
	"time"
	"transaction/internal/entity"
)

type Service interface {
	GetUserBalance(userID int) ([]entity.CardBalance, error)
	CreateInvoice(cardNumber int, currency entity.Currency, amount float64) (*entity.Transaction, error)
	CreateWithdraw(cardNumber int, currency entity.Currency, amount float64) (*entity.Transaction, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s service) GetUserBalance(userID int) ([]entity.CardBalance, error) {
	return s.repository.CountUserBalance(userID)
}

func (s service) CreateInvoice(
	cardNumber int,
	currency entity.Currency,
	amount float64) (*entity.Transaction, error) {

	trans := entity.Transaction{

		CardNumber: cardNumber,
		Type:       entity.Invoice,
		Currency:   currency,
		Amount:     amount,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Status:     entity.Created,
	}
	card, err := s.repository.GetCard(cardNumber)
	if err != nil {
		return nil, err
	}
	if card.Currency != currency {
		return nil, errors.New("invalid currency")
	}
	err = s.repository.CreateTransaction(&trans)
	return &trans, err
}

var cardWithdrawLock map[int]sync.Mutex

func getMutex(cardNumber int) {
	// todo
}

func (s service) CreateWithdraw(
	cardNumber int,
	currency entity.Currency,
	amount float64) (*entity.Transaction, error) {

	trans := entity.Transaction{
		CardNumber: cardNumber,
		Type:       entity.Withdraw,
		Currency:   currency,
		Amount:     amount,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Status:     entity.Created,
	}
	// TODO lock card for withdraw
	card, err := s.repository.GetCard(cardNumber)
	if err != nil {
		return nil, err
	}
	if card.Currency != currency {
		return nil, errors.New("invalid currency")
	}

	cardBalance, err := s.repository.CountCardBalance(cardNumber)
	if err != nil {
		return nil, err
	}
	if cardBalance.ActualBalance-cardBalance.PendingWithdraw-amount < 0 {
		return nil, errors.New("not enough money")
	}

	err = s.repository.CreateTransaction(&trans)
	return &trans, err
}

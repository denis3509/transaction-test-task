package transaction

import (
	"transaction/internal/entity"
	"transaction/pkg/db"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Repository interface {
	CountUserBalance(UserID int) ([]entity.CardBalance, error)
	CountCardBalance(cardNumber int) (entity.CardBalance, error)
	CreateTransaction(trans *entity.Transaction) error
	UpdateTransaction(trans *entity.Transaction, fields ...string) error
	GetCard(cardNumber int) (entity.Card, error)
}
 
type repository struct {
	db *db.DB
	// logger log.Logger
}

func NewRepository(db *db.DB) Repository {
	return &repository{db}
}
func (r *repository) CountUserBalance(UserID int) ([]entity.CardBalance, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repository) CountCardBalance(cardNumber int) (entity.CardBalance,error) {
	panic("not implemented") // TODO: Implement
}

func (r *repository) CreateTransaction(trans *entity.Transaction) error {
	err := r.db.Model(trans).Insert()
	return err
}

func (r *repository) UpdateTransaction(trans *entity.Transaction, fields ...string) error {
	return r.db.Model(trans).Update(fields...)
}

func (r *repository) GetCard(cardNumber int) (entity.Card, error) {
	card := entity.Card{}
	err := r.db.Select("id", "name").
			From("users").
			Where(dbx.HashExp{"id": 100}).
			One(&card)
	return card, err 
}

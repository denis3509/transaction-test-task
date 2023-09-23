package transaction

import (
	"transaction/internal/entity"
	"transaction/pkg/db"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Repository interface {
	CountUserBalance(UserID int) ([]entity.CardBalance, error)
	CountCardBalance(cardNumber int) (*entity.CardBalance, error)
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
	var cards []entity.CardBalance
	sql := `
	select card_number,
       s.total_success_invoice - s.total_success_withdraw as actual_balance,
       pending_invoice,
       pending_withdraw
       from
(select card_number,
       sum(case when status=2 and type= 1
           then amount
           else   0
           end) as total_success_invoice,
       sum(case when status=2 and type= 2
           then amount
           else   0
           end) as total_success_withdraw,
       sum(case when status=1 and type= 1
           then amount
           else   0
           end) as pending_invoice,
       sum(case when status=1 and type= 2
           then amount
           else   0
           end) as pending_withdraw
from transaction
                 left join card
on card.number = transaction.card_number
                 where card.user_id = {:user_id}
group by card_number) as s	
`
	q := r.db.NewQuery(sql)
	q.Bind(dbx.Params{"user_id": UserID})

	err := q.All(&cards)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (r *repository) CountCardBalance(cardNumber int) (*entity.CardBalance, error) {
	var card entity.CardBalance
	sql := `
	select card_number,
       s.total_success_invoice - s.total_success_withdraw as actual_balance,
       pending_invoice,
       pending_withdraw
       from
(select card_number,
       sum(case when status=2 and type= 1
           then amount
           else   0
           end) as total_success_invoice,
       sum(case when status=2 and type= 2
           then amount
           else   0
           end) as total_success_withdraw,
       sum(case when status=1 and type= 1
           then amount
           else   0
           end) as pending_invoice,
       sum(case when status=1 and type= 2
           then amount
           else   0
           end) as pending_withdraw
from transaction
                 left join card
on card.number = transaction.card_number
                 where card_number = {:card_number}
group by card_number) as s	
`
	q := r.db.NewQuery(sql)
	q.Bind(dbx.Params{"card_number": cardNumber})

	err := q.One(&card)
	if err != nil {
		return nil, err
	}
	return &card, nil
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
	err := r.db.Select().
		From("card").
		Where(dbx.HashExp{"number": cardNumber}).
		One(&card)
	return card, err
}

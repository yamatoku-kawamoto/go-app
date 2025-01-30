package repository

import (
	"goapp/internal/core/entity"
	"goapp/internal/repository/database"
)

type Transaction struct {
	db *database.DB
}

func (t *Transaction) FindByID(id int64) (*entity.Transaction, error) {
	return nil, nil
}

func (t *Transaction) GetByCustomerID(id int64) ([]*entity.Transaction, error) {
	return nil, nil
}

func (t *Transaction) GetByCustomerName(name string) ([]*entity.Transaction, error) {
	return nil, nil
}

func (t *Transaction) GetByCustomerEmail(email string) ([]*entity.Transaction, error) {
	return nil, nil
}

func (t *Transaction) GetByCustomerPhone(phone string) ([]*entity.Transaction, error) {
	return nil, nil
}

func (t *Transaction) GetByCustomerFuriganaName(furiganaName string) ([]*entity.Transaction, error) {
	return nil, nil
}

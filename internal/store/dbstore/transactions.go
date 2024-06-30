package dbstore

import (
	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"gorm.io/gorm"
)

type TransactionStore struct {
	db *gorm.DB
}

func NewTransactionStore() *TransactionStore {
	return &TransactionStore{
		db: initializers.DB,
	}
}

func (s *TransactionStore) CreateTransaction(transaction *store.Transaction) error {
	return s.db.Create(transaction).Error
}

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

func (s *TransactionStore) UpdateTransactionState(checkoutId string, state string) error {
	var transaction store.Transaction
	err := s.db.
		Model(&transaction).
		Where("checkout_request_id", checkoutId).
		Update("payment_status", state).Error
	return err
}

func (s *TransactionStore) GetTransactionByUserId() *store.Transaction {
	var transaction store.Transaction
	return &transaction
}

func (s *TransactionStore) SaveTransaction(transaction *store.Transaction) error {
	return s.db.Save(&transaction).Error
}

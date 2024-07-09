package dbstore

import (
	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"gorm.io/gorm"
)

type TransactionStore struct {
	db     *gorm.DB
	UserID uint
}

func NewTransactionStore(userId uint) *TransactionStore {
	return &TransactionStore{
		UserID: userId,
		db:     initializers.DB,
	}
}

func (s *TransactionStore) GetTransactionByUserId() *store.Transaction {
	var transaction store.Transaction
	transaction.UserID = s.UserID
	s.db.Preload("Orders").Where("user_id = ?", s.UserID).First(&transaction)

	return &transaction
}

func (s *TransactionStore) SaveTransaction(transaction *store.Transaction) error {
	return s.db.Save(&transaction).Error
}

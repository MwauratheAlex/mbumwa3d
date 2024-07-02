package dbstore

import (
	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"gorm.io/gorm"
)

type OrderStore struct {
	db *gorm.DB
}

func NewOrderStore() *OrderStore {
	return &OrderStore{
		db: initializers.DB,
	}
}

func (s *OrderStore) CreateOrder(order *store.Order) error {
	return s.db.Create(order).Error
}

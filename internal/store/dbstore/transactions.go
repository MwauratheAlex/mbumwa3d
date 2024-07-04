package dbstore

import (
	"fmt"

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

func (s *OrderStore) GetNotCompleted(userID uint) []store.Order {
	var orders []store.Order

	s.db.Preload("File").Where("user_id = ? AND status != ?",
		userID,
		fmt.Sprint(store.Completed),
	).Find(&orders)

	return orders
}

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
	query := "user_id = ? AND status != ?"

	s.db.Preload("File").Where(query,
		userID,
		fmt.Sprint(store.Completed),
	).Find(&orders)

	return orders
}

func (s *OrderStore) GetCompleted(userID uint) []store.Order {
	var orders []store.Order
	query := "user_id = ? AND status = ?"

	s.db.Preload("File").Where(query,
		userID,
		fmt.Sprint(store.Completed),
	).Find(&orders)

	return orders
}

func (s *OrderStore) GetPrintAvailable(userID uint) []store.Order {
	var orders []store.Order

	s.db.Preload("File").Find(&orders)

	return orders
}

func (s *OrderStore) GetPrintActive(userID uint) []store.Order {
	var orders []store.Order

	s.db.Preload("File").Find(&orders)

	return orders
}

func (s *OrderStore) GetPrintCompleted(userID uint) []store.Order {
	var orders []store.Order

	s.db.Preload("File").Find(&orders)

	return orders
}

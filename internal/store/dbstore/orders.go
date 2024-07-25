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

func (s *OrderStore) GetByID(orderID uint) (*store.Order, error) {

	order := store.Order{}
	err := s.db.Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, err
}

func (s *OrderStore) CreateOrder(order *store.Order) error {
	return s.db.Create(order).Error
}

func (s *OrderStore) Save(order *store.Order) error {
	return s.db.Save(&order).Error
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

func (s *OrderStore) GetPrintAvailable() []store.Order {
	var orders []store.Order

	s.db.Preload("File").Where("print_status = ?", fmt.Sprint(store.Available)).Find(&orders)

	return orders
}

func (s *OrderStore) GetPrintActive(printerID uint) []store.Order {
	var orders []store.Order

	s.db.Preload("File").Where(
		"print_status = ? AND printer_id = ?",
		store.Selected, printerID,
	).Find(&orders)

	return orders
}

func (s *OrderStore) GetPrintCompleted(printerID uint) []store.Order {
	var orders []store.Order

	s.db.Debug().Preload("File").
		Where(
			"print_status = ? AND printer_id = ?",
			fmt.Sprint(store.Completed),
			printerID,
		).Find(&orders)

	return orders
}
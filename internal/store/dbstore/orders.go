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

func (s *OrderStore) GetByID(orderID uint) (*store.Order, error) {

	order := store.Order{}
	err := s.db.Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, err
}

func (s *OrderStore) CreatePrintConfig(printConfig *store.PrintConfig) error {
	return s.db.Create(printConfig).Error
}

func (s *OrderStore) CreateOrder(order *store.Order) error {
	return s.db.Create(order).Error
}

func (s *OrderStore) Save(order *store.Order) error {
	return s.db.Save(&order).Error
}

func (s *OrderStore) Delete(userID, orderID uint) error {
	return s.db.
		Where("user_id = ?", userID).
		Delete(&store.Order{ID: orderID}).Error
}

func (s *OrderStore) GetAvailable() ([]store.Order, error) {
	var orders []store.Order
	err := s.db.Preload("PrintConfig").
		Where("status = ?", store.Reviewing.String()).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderStore) GetPrintOrders(
	printerID uint, orderStatus string) ([]store.Order, error) {
	var orders []store.Order

	err := s.db.Preload("PrintConfig").Where(
		"status = ? AND printer_id = ?",
		orderStatus, printerID,
	).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *OrderStore) GetNotCompleted(userID uint) []store.Order {
	var orders []store.Order
	query := "user_id = ? AND status != ?"

	s.db.Preload("PrintConfig").Where(query,
		userID,
		store.Completed.String(),
	).Find(&orders)

	return orders
}

func (s *OrderStore) GetCompleted(userID uint) []store.Order {
	var orders []store.Order
	query := "user_id = ? AND status = ?"

	s.db.Preload("PrintConfig").Where(query,
		userID,
		store.Completed.String(),
	).Find(&orders)

	return orders
}

package dbstore

import (
	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore() *UserStore {
	return &UserStore{
		db: initializers.DB,
	}
}

func (s *UserStore) CreateUser(email string, password string, hasPrinter bool) error {
	return s.db.Create(&store.User{
		Email:      email,
		HasPrinter: hasPrinter,
	}).Error
}

func (s *UserStore) GetUser(email string) (*store.User, error) {
	user := store.User{}
	err := s.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (s *UserStore) GetUserById(id uint) (*store.User, error) {
	user := store.User{}
	err := s.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (s *UserStore) GetOrCreate(user *store.User) (*store.User, error) {
	err := s.db.Where("email = ?", user.Email).FirstOrCreate(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserStore) GetOrder(
	orderID, userID uint) (*store.Order, error) {

	order := store.Order{}
	err := s.db.
		Preload("PrintConfig").
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order).Error

	if err != nil {
		return nil, err
	}
	return &order, err
}

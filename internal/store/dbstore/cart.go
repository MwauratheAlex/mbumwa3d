package dbstore

import (
	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"gorm.io/gorm"
)

type CartStore struct {
	UserID uint
	db     *gorm.DB
}

func NewCartStore(userID uint) *CartStore {
	return &CartStore{
		db:     initializers.DB,
		UserID: userID,
	}
}

func (s *CartStore) Create() {}

func (s *CartStore) GetCartByUserId() *store.Cart {
	var cart store.Cart
	cart.UserID = s.UserID
	s.db.Debug().Preload("Orders").Where("user_id = ?", s.UserID).FirstOrCreate(&cart)

	return &cart
}

func (s *CartStore) SaveCart(cart *store.Cart) error {
	return s.db.Save(&cart).Error
}

func (s *CartStore) AddItemToCart()      {}
func (s *CartStore) RemoveItemFromCart() {}

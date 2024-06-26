package dbstore

import (
	"fmt"

	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserStore) CreateUser(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(email), 10)
	if err != nil {
		return err
	}
	fmt.Println(hashedPassword)
	return nil
}

func (s *UserStore) GetUser(email string) (*store.User, error) {
	user := store.User{}

	return &user, nil
}

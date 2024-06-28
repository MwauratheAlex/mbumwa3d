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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	fmt.Println(email, password, string(hashedPassword))
	return s.db.Create(&store.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
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

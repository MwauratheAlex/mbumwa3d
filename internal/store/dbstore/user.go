package dbstore

import "github.com/mwaurathealex/mbumwa3d/internal/store"

type UserStore struct {
	db           string
	passwordHash string
}

type NewUserStoreParams struct {
	DB           string
	PasswordHash string
}

func NewUserStore(params NewUserStoreParams) *UserStore {
	return &UserStore{
		db:           params.DB,
		passwordHash: params.PasswordHash,
	}
}

func (s *UserStore) CreateUser(email string, password string) error {
	return nil
}

func (s *UserStore) GetUser(email string) (*store.User, error) {
	user := store.User{}

	return &user, nil
}

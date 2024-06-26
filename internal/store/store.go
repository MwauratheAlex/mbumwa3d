package store

type User struct {
	ID             uint
	Email          string
	HashedPassword string
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}

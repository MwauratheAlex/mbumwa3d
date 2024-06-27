package store

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Email        string    `gorm:"type:citext;unique;not null"`
	PasswordHash string    `gorm:"type:varchar(255)"`
	InsertedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}

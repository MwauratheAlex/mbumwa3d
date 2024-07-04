package store

import (
	"mime/multipart"
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Email        string `gorm:"type:citext;unique;not null"`
	PasswordHash string `gorm:"type:varchar(255)"`
	HasPrinter   bool
	InsertedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type UserStore interface {
	CreateUser(email string, password string, hasPrinter bool) error
	GetUser(email string) (*User, error)
	GetUserById(id uint) (*User, error)
}

type File struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	UserID     uint
	User       User      `gorm:"foreignKey:UserID"`
	InsertedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	LocalPath  string
	FileName   string
	Technology string

	Color string
}

type FileStore interface {
	SaveToDisk(file multipart.File, filename string) (string, error)
}

type OrderState int

const (
	Reviewing OrderState = iota
	Processing
	Shipping
	Completed
)

func (os OrderState) String() string {
	return [...]string{"Reviewing", "Processing", "Shipping", "Completed"}[os]
}

type Order struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	UserID     uint
	FileID     uint
	User       User      `gorm:"foreignKey:UserID"`
	File       File      `gorm:"foreignKey:FileID"`
	InsertedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	BuildTime       uint
	Quantity        string
	Price           float64
	Phone           string
	PaymentComplete bool
	Status          string
}

type OrderStore interface {
	Createorder(*Order) error
}

type Cart struct {
	ID     uint    `gorm:"primaryKey;autoIncrement"`
	UserID uint    `gorm:"index"`
	Orders []Order `gorm:"many2many:cart_orders;"`
}

type CartStore interface {
	Create(string)
	GetCartByUserId(string) *Cart
	AddItem()
	RemoveItem()
}

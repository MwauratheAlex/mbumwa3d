package store

import (
	"mime/multipart"
	"time"
)

type State int

const (
	AwaitingPayment State = iota
	Reviewing
	Printing
	Shipping
	Completed
)

func (os State) String() string {
	return [...]string{
		"AwaitingPayment",
		"Reviewing",
		"Printing",
		"Shipping",
		"Completed",
	}[os]
}

type User struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Email      string `gorm:"type:citext;unique;not null"`
	Name       string
	PhotoUrl   string
	Orders     []Order
	HasPrinter bool
	InsertedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type UserStore interface {
	CreateUser(email string, password string, hasPrinter bool) error
	GetUser(email string) (*User, error)
	GetUserById(id uint) (*User, error)
	GetOrCreate(user *User) (*User, error)
}

type PrintConfig struct {
	ID         uint
	Technology string
	Material   string
	Color      string
	Quantity   int
	FileID     string // SaveToDisk name OR FileID in cloudflare
	User       User
	UserID     uint
	FileVolume float64
	Price      float64
}

type SummaryModalParams struct {
	IsLoggedInUser bool
	PrintContif    PrintConfig
}

type FileStore interface {
	SaveToDisk(file multipart.File, filename string) (string, error)
}

type Order struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	UserID        uint `gorm:"index"`
	User          User `gorm:"foreignKey:UserID"`
	PrinterID     *uint
	Printer       User `gorm:"foreignKey:PrinterID"`
	PrintConfigID uint
	PrintConfig   PrintConfig `gorm:"foreignKey:PrintConfigID"`

	BuildTime         uint
	Price             float64
	PaymentComplete   bool
	Status            string
	CheckoutRequestId string

	InsertedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type OrderStore interface {
	Createorder(*Order) error
}

type Cart struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	UserID        uint `gorm:"index"`
	TransactionID uint
	Transaction   *Transaction
}

type Transaction struct {
	ID                uint    `gorm:"primaryKey;autoIncrement"`
	UserID            uint    `gorm:"index"`
	Orders            []Order `gorm:"many2many:transaction_orders;"`
	PaymentStatus     string
	CheckoutRequestId string
	Phone             string
}
type TransactionStore interface {
	GetTransactionByUserId() *Transaction
	SaveTransaction(*Transaction) error
}

type CartStore interface {
	Create(string)
	GetCartByUserId(string) *Cart
	AddItem()
	RemoveItem()
}

package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := getConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn)
		panic(fmt.Sprintf("Error connecting to database: %s", err))
	} else {
		fmt.Println("Connected to db successfully")
		DB = db
	}
}

func getConnectionString() string {

	if os.Getenv("env") == "production" {
		return os.Getenv("DATABASE_URL")
	}
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_NAME")
	hostname := os.Getenv("PG_HOST")

	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		user, password, hostname, dbname)
}

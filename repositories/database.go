package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

const (
	host = "localhost"
	port = 5432
	username = "jlhe"
	databaseName = "simple_api_demo"
)

func ConnectToDatabase() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, username, databaseName)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	database = db
}

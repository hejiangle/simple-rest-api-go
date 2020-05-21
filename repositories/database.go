package repositories

import (
	"../dto"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

const (
	host         = "localhost"
	port         = 5432
	username     = "jlhe"
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

func Migration() {
	database.DropTableIfExists(&dto.TodoItem{})
	database.AutoMigrate(&dto.TodoItem{})
	initData := initToDoItems()
	for _, item := range initData {
		database.Create(&item)
	}
}

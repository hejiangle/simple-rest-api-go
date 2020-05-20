package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var database *sql.DB

const (
	host = "localhost"
	port = 5432
	username = "jlhe"
	databaseName = "simple_api_demo"
)

func ConnectToDatabase() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, username, databaseName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	database = db
}

func GetDatabase() *sql.DB {
	return database
}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	username = "jlhe"
	database = "simple_api_demo"
)

func connectToDatabase() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, username, database)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

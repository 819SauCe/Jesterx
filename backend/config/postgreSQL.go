package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var DB *sql.DB

func ConnectPostgres() {
	var err error
	DB, err = sql.Open("postgres", PostgresUri)
	if err != nil {
		log.Fatal("Error on Postgres: ", err)
	}

	if err := DB.Ping(); err != nil {
		fmt.Println("Error on Postgres: ")
		fmt.Println("Try again in 5 seconds")
		time.Sleep(time.Second * 5)
		ConnectPostgres()
	}

	log.Println("Connected to Postgres!")
}

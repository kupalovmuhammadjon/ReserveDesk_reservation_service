package main

import (
	"fmt"
	"log"
	"reservation_service/storage/postgres"
)

func main() {
	fmt.Println("👋")
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("postgres connection error: ", err)
	}
	defer db.Close()
	fmt.Println("👋")
}
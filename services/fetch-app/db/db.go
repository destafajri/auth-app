package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
  

func DBcon() *sql.DB{
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Println("Error connection")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error ping")
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}

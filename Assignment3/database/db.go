package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// koneksikan database nya
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	DB, err = sql.Open("postgres", config)

	if err != nil {
		return nil, err
	}

	log.Println("Koneksi database aman.")

	createTableSQL := `
		CREATE TABLE IF NOT EXISTS weather (
			id SERIAL PRIMARY KEY,
			water INT,
			wind INT,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err = DB.Exec(createTableSQL)

	if err != nil {
		return nil, err
	}

	return DB, nil
}

func GetDB() *sql.DB {
	return DB
}

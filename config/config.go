package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Configurar o pool de conexões
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(time.Minute * 5)

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	DB = db
	log.Println("Successfully connected to the database")
}

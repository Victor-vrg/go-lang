package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	// Usar a DATABASE_URL fornecida pelo Railway
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Abrir a conexão com o banco de dados
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Testar a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	DB = db
	log.Println("Successfully connected to the database")
}

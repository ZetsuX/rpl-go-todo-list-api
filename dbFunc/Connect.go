package dbFunc

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://"+getEnv("POSTGRE_USER")+":"+getEnv("POSTGRE_PASSWORD")+"@localhost:5432/oprecrpl")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
		return ""
	}

	return os.Getenv(key)
}

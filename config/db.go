package config

import (
	"database/sql"
	"fmt"
)

func SetupDB() *sql.DB {
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable")
}

package services

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	return db, err
}

package main

import (
	"database/sql"
	"fmt"
	"github.com/atsuof/sampleapi/api"
	"log"
	"net/http"
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

func main() {

	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		closeErr := db.Close()
		if closeErr != nil {
			panic(fmt.Errorf("an error occurred when closing db connections.:%w", closeErr))
		}
	}()
	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package repositories_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

var testDB *sql.DB

func setUp() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbName := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)
	db, dbOpenError := sql.Open("mysql", dbConn)
	if dbOpenError != nil {
		return dbOpenError
	}
	testDB = db
	return nil
}

func teardown() {
	err := testDB.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	// Setup code here, if needed
	if err := setUp(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//code := m.Run()
	code := m.Run()
	// Teardown code here, if needed
	teardown()
	os.Exit(code)
}

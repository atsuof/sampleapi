package main

import (
	"database/sql"
	"fmt"
	"github.com/atsuof/sampleapi/controllers"
	"github.com/atsuof/sampleapi/services"
	"github.com/gorilla/mux"
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
	service := services.NewMyAppService(db)
	controller := controllers.NewMyAppControllers(service)

	r := mux.NewRouter()

	r.HandleFunc("/hello", controller.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", controller.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", controller.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+$}", controller.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", controller.PostArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/comment", controller.PostArticleCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

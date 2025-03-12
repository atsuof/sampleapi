package api

import (
	"database/sql"
	"github.com/atsuof/sampleapi/controllers"
	"github.com/atsuof/sampleapi/services"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(db *sql.DB) *mux.Router {
	service := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(service)
	cCon := controllers.NewCommentController(service)
	r := mux.Router{}
	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+$}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/comment", cCon.PostArticleCommentHandler).Methods(http.MethodPost)
	return &r
}

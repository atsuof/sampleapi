package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/atsuof/sampleapi/models"
	"github.com/atsuof/sampleapi/services"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type MyAppControllers struct {
	service *services.MyAppService
}

func NewMyAppControllers(service *services.MyAppService) *MyAppControllers {
	return &MyAppControllers{service: service}
}

func (s *MyAppControllers) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// PostArticleHandler function handles POST requests for posting an article.
// It handles request to "/article"
func (s *MyAppControllers) PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var article models.Article
	if decErr := json.NewDecoder(req.Body).Decode(&article); decErr != nil {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
	}

	registered, err := s.service.PostArticleService(article)
	if err != nil {
		http.Error(w, "registration the Article is failed", http.StatusInternalServerError)
		return
	}

	if encErr := json.NewEncoder(w).Encode(registered); encErr != nil {
		http.Error(w, "fail to create response body\n", http.StatusBadRequest)
	}
}

// ArticleListHandler handles GET requests for fetching a list of articles, optionally filtered by page query parameter.
// It handles request to "/article/list"
func (s *MyAppControllers) ArticleListHandler(w http.ResponseWriter, req *http.Request) {

	p, _ := req.URL.Query()["page"]
	page := 1
	if len(p) > 0 {
		tmpPage, err := strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		page = tmpPage
	}

	articles, err := s.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "Failed to get articles", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&articles); err != nil {
		http.Error(w, "failed to marshal json data", http.StatusBadRequest)
	}
}

func (s *MyAppControllers) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid Query parameter", http.StatusBadRequest)
		return
	}

	article, err := s.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "failed to get article", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(&article); err != nil {
		http.Error(w, "failed to marshal json data", http.StatusBadRequest)
	}
}

func (s *MyAppControllers) PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {

	id, ok := req.URL.Query()["id"]
	if !ok {
		http.Error(w, "Failed to like the article", http.StatusInternalServerError)
	}
	articleID, err := strconv.Atoi(id[0])

	err = s.service.PostNiceService(articleID)
	if err != nil {
		http.Error(w, "Failed to like the article", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	_, writeErr := fmt.Fprintln(w, "Success to like the article")
	if writeErr != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (s *MyAppControllers) PostArticleCommentHandler(w http.ResponseWriter, req *http.Request) {

	var comment models.Comment
	if decErr := json.NewDecoder(req.Body).Decode(&comment); decErr != nil {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
	}

	registered, err := s.service.PostCommentService(comment)

	if err != nil {
		http.Error(w, "Failed to register the comment", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(&registered); err != nil {
		http.Error(w, "failed to marshal json data", http.StatusBadRequest)
	}

}

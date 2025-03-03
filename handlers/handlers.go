package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/atsuof/sampleapi/models"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// PostArticleHandler function handles POST requests for posting an article.
// It handles request to "/article"
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "Posting Articl..\n")

	var article1 models.Article
	if decErr := json.NewDecoder(req.Body).Decode(&article1); decErr != nil {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
	}

	if encErr := json.NewEncoder(w).Encode(article1); encErr != nil {
		http.Error(w, "fail to create response body\n", http.StatusBadRequest)
	}
}

// ArticleListHandler handles GET requests for fetching a list of articles, optionally filtered by page query parameter.
// It handles request to "/article/list"
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {

	p, _ := req.URL.Query()["page"]
	//if !ok {
	//	http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	//	return
	//}

	page := 1
	if len(p) > 0 {
		tmpPage, err := strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		page = tmpPage
	}

	log.Println(page)

	articles := []models.Article{models.Article1, models.Article2}
	if err := json.NewEncoder(w).Encode(&articles); err != nil {
		http.Error(w, "failed to marshal json data", http.StatusBadRequest)
	}
}

func returnJsonDatas() ([]byte, error) {
	articles := []models.Article{models.Article1, models.Article2}
	jsonDates, err := json.Marshal(articles)
	if err != nil {
		return nil, err
	}
	return jsonDates, nil
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid Query parameter", http.StatusBadRequest)
		return
	}
	//resString := fmt.Sprintf("Article No:%d\n", articleID)
	//io.WriteString(w, resString)
	fmt.Println(articleID)
	jsonData, err := json.Marshal(models.Article1)
	if err != nil {
		http.Error(w, "failed to marshal json data", http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "Posting Nice...\n")

	jsonData, err := json.Marshal(models.Article1)
	if err != nil {
		http.Error(w, "failed to marshal json data", http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func PostArticleCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
	jsonData, err := json.Marshal(models.Comment1)
	if err != nil {
		http.Error(w, "failed to marshal json data", http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

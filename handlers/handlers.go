package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Articl..\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List..\n")
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid Query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No:%d\n", articleID)
	io.WriteString(w, resString)
}

func PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostArticleCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}

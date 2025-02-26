package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/atsuof/reponame/models"
	"github.com/gorilla/mux"
	"io"
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

	contentsLen, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "Invalid Content-Length\n", http.StatusBadRequest)
		return
	}
	requestBuf := make([]byte, contentsLen)

	if _, err := req.Body.Read(requestBuf); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var reqArticle models.Article

	if unmarshalError := json.Unmarshal(requestBuf, &reqArticle); unmarshalError != nil {
		http.Error(w, "failed to decode json data \n", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(reqArticle)
	if err != nil {
		http.Error(w, "faild to encode json", http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

// ArticleListHandler handles GET requests for fetching a list of articles, optionally filtered by page query parameter.
// It handles request to "/article/list"
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {

	if p, ok := req.URL.Query()["page"]; ok && len(p) > 0 {
		page, err := strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		fmt.Println(page)
		//io.WriteString(w, fmt.Sprintf("Article List page%d\n", page))
		jsonDates, err := returnJsonDatas()
		if err != nil {
			http.Error(w, "failed to marshal json data", http.StatusBadRequest)
		} else {
			w.Write(jsonDates)
		}

	} else {
		jsonDates, err := returnJsonDatas()
		if err != nil {
			http.Error(w, "failed to marshal json data", http.StatusBadRequest)
		} else {
			w.Write(jsonDates)
		}
	}

	io.WriteString(w, "Article List..\n")
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

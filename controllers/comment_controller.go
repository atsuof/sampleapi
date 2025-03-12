package controllers

import (
	"encoding/json"
	"github.com/atsuof/sampleapi/controllers/services"
	"github.com/atsuof/sampleapi/models"
	"net/http"
)

type CommentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) *CommentController {
	return &CommentController{service: service}
}

func (s *CommentController) PostArticleCommentHandler(w http.ResponseWriter, req *http.Request) {

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

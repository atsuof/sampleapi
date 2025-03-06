package services

import (
	"github.com/atsuof/sampleapi/models"
	"github.com/atsuof/sampleapi/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	registered, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return registered, nil
}

package services

import (
	"database/sql"
	"errors"
	"github.com/atsuof/sampleapi/apperrors"
	"github.com/atsuof/sampleapi/models"
	"github.com/atsuof/sampleapi/repositories"
)

type MyAppService struct {
	db *sql.DB
}

func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}

func (s *MyAppService) GetArticleService(articleiId int) (models.Article, error) {

	article, err := repositories.SelectArticleDetail(s.db, articleiId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	comments, err := repositories.SelectCommentList(s.db, articleiId)
	if err != nil {
		return models.Article{}, err
	}
	article.Comments = comments
	return article, err
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {

	registered, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "filed to record data")
		return models.Article{}, err
	}
	return registered, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {

	articles, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (s *MyAppService) PostNiceService(articleID int) error {

	err := repositories.UpdateNiceNum(s.db, articleID)
	if err != nil {
		return err
	}
	return nil
}

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {

	registered, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return registered, nil
}

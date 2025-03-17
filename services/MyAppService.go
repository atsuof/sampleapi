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

	type articleResult struct {
		article models.Article
		err     error
	}
	artChan := make(chan articleResult)
	defer close(artChan)
	go func(ch chan<- articleResult, db *sql.DB, articleId int) {
		article, err := repositories.SelectArticleDetail(s.db, articleId)
		ch <- articleResult{article: article, err: err}
	}(artChan, s.db, articleiId)

	type commentResult struct {
		comments *[]models.Comment
		err      error
	}

	comChan := make(chan commentResult)
	defer close(comChan)
	go func(ch chan<- commentResult, db *sql.DB, articleId int) {
		comments, err := repositories.SelectCommentList(s.db, articleId)
		ch <- commentResult{
			comments: &comments,
			err:      err,
		}
	}(comChan, s.db, articleiId)

	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	for i := 0; i < 2; i++ {
		select {
		case res := <-artChan:
			article, articleGetErr = res.article, res.err
		case res := <-comChan:

			commentList, commentGetErr = *res.comments, res.err
		}
	}
	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			articleGetErr = apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, articleGetErr
		}
		articleGetErr = apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, articleGetErr
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	article.Comments = commentList
	return article, nil
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
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
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

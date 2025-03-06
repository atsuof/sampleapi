package services

import (
	"github.com/atsuof/sampleapi/models"
	"github.com/atsuof/sampleapi/repositories"
)

func GetArticleService(articleiId int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleiId)
	if err != nil {
		return models.Article{}, err
	}

	comments, err := repositories.SelectCommentList(db, articleiId)
	if err != nil {
		return models.Article{}, err
	}
	article.Comments = comments
	return article, err
}

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	registered, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}
	return registered, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articles, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func PostNiceService(articleID int) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, articleID)
	if err != nil {
		return err
	}
	return nil
}

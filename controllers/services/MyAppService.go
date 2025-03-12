package services

import "github.com/atsuof/sampleapi/models"

type ArticleService interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(articleId int) error
}

type CommentService interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}

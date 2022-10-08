package services

import "github.com/yanachuwan9sm/myapi-tutorial/models"

// /article 関連を引き受けるサービス
type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

// /comment 関連を引き受けるサービス
type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}

package services

import (
	"database/sql"
	"errors"

	"github.com/yanachuwan9sm/myapi-tutorial/apperrors"
	"github.com/yanachuwan9sm/myapi-tutorial/models"
	"github.com/yanachuwan9sm/myapi-tutorial/repositories"
)

// ArticleDetailHandlerで使うことを想定したサービス
// 指定 ID の記事情報を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {

	// 1. repositories層の関数SelectArticleDetailで記事の詳細を取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		// ErrNoRows は、QueryRow メソッドが全く列を取得できなかったときに
		// Scan メソッドが呼ばれたときに返却される
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	// 2. repositories層の関数SelectCommentListでコメント一覧を取得
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	// 3. 2で得たコメント一覧を、1で得たArticle構造体に紐付ける
	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

// PostArticleHandlerで使うことを想定したサービス
// 引数の情報をもとに記事データをDB内に挿入し、結果を返却
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {

	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

// ArticleListHandlerで使うことを想定したサービス
// 指定されたページの記事一覧をデータベースから取得し、取得した値を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {

	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	// db.QueryRow メソッドで取得結果 0 件の場合は ErrNoRows エラーが返却されるが、
	// しかし db.Query メソッドの場合は、取得結果が0件であるという正常応答が返ってきてエラーとならない
	// -> 得られた結果が0件かどうかをその場で判定してエラーを返す仕組みを追加実装する

	// SelectArticleList関数から取得した記事の長さが0だった場合
	if len(articleList) == 0 {
		err = apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

// PostNiceHandlerで使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {

	err := repositories.UpdateNiceNum(s.db, article.ID)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist taget data")
			return models.Article{}, err
		}

		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")

		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}

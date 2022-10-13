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

	// Article型とerror型を同時に扱う構造体 (とあるチャネルで送信できる型は一つだけ)
	type articleResult struct {
		article models.Article
		err     error
	}

	// articleResult型のチャネルを定義
	articleChan := make(chan articleResult)
	defer close(articleChan)

	// repositories層の関数SelectArticleDetailで記事の詳細を取得
	go func(ch chan<- articleResult) {
		// articleChanを通じて、SelectArticleDetail関数の結果を送信
		article, err := repositories.SelectArticleDetail(s.db, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan)

	// Comment型のスライスとerror型を同時に扱う構造体
	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}

	// commentResult型のチャネルを定義
	commentChan := make(chan commentResult)
	defer close(commentChan)

	// repositories層の関数SelectCommentListでコメント一覧を取得
	go func(ch chan<- commentResult) {
		// commentChanを通じて、SelectCommentList関数の結果を送信
		commentList, err := repositories.SelectCommentList(s.db, articleID)
		ch <- commentResult{commentList: &commentList, err: err}
	}(commentChan)

	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		}
	}

	if articleGetErr != nil {
		// ErrNoRows は、QueryRow メソッドが全く列を取得できなかったときに
		// Scan メソッドが呼ばれたときに返却される
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			articleGetErr = apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, articleGetErr
		}
		articleGetErr = apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, articleGetErr
	}

	if commentGetErr != nil {
		commentGetErr = apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, commentGetErr
	}

	// コメント一覧をArticle構造体に紐付ける
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

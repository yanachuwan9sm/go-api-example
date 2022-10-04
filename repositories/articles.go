package repositories

import (
	"database/sql"

	"github.com/yanachuwan9sm/myapi-tutorial/models"
)

const (
	articleNumPerPage = 5
)

// 新規投稿をデータベースに insert する関数
// -> データベースに保存した記事内容と、発生したエラーを返り値にする
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {

	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)

	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

// 変数 page で指定されたページに表示する投稿一覧をデータベースから取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {

	const sqlStr = `
	select article_id, title, contents, username, nice
	from articles
	limit ? offset ?;
	`
	// 指定された記事データをデータベースから取得
	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// rowsからデータを取り出し models.Article構造体に格納

	//models.Article型のスライスarticleArrayを生成
	articleArray := make([]models.Article, 0)

	// 取り出したデータを models.Article 構造体のスライス []models.Article に詰めて返す処理
	for rows.Next() {

		var article models.Article

		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

// 投稿IDを指定して、記事データを取得する関数
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime

	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

// いいねの数を update する関数
// -> 発生したエラーを返り値にする
func UpdateNiceNum(db *sql.DB, articleID int) error {

	// 指定された ID の記事のいいね数を+1 するようにデータベースの中身を更新する処理

	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// 現在のいいね数を取得するクエリの実行し変数 nicenum に取得した値を読み込む

	const sqlGetNice = ` select nice
			from articles
			where article_id = ?;
		`
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var nicenum int
	err = row.Scan(nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	// いいね数を +1 する処理
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

package models

import "time"

// Comment構造体
type Comment struct {
	CommentID int       `json:"comment_id"` // コメントID - ブログサービスに投稿される全コメントに振られる連番
	ArticleID int       `json:"article_id"` // コメント対象となった記事 ID - どの記事に対するコメントなのかを明示するために、対象記事のIDを含める
	Message   string    `json:"message"`    // コメント本文
	CreatedAt time.Time `json:"created_at"` // 投稿日時
}

// 記事構造体
type Article struct {
	ID          int       `json:"article_id"` // 記事 ID: ブログサービスに投稿される全記事に振られる連番
	Title       string    `json:"title"`      // 記事タイトル
	Contents    string    `json:"contents"`   // 本文記事
	UserName    string    `json:"user_name"`  // ユーザー名
	NiceNum     int       `json:"nice"`       // いいね数
	CommentList []Comment `json:"comments"`   // その記事に紐ついたコメントをスライス形式で格納
	CreatedAt   time.Time `json:"created_at"` //投稿日時
}

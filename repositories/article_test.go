package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yanachuwan9sm/myapi-tutorial/models"
	"github.com/yanachuwan9sm/myapi-tutorial/repositories"
)

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// 期待値を定義
	expected := models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  2,
	}

	// テスト対象となる関数を実行
	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		// SelectArticleDetailがうまくいかなくてそもそも返り値gotが得られていないなら
		// この後の期待する出力と実際の出力の比較が不可能なのでテスト続行不可
		t.Fatal(err)
	}

	// 評価

	// 記事IDが同じかどうか比較
	if got.ID != expected.ID {
		t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
	}
	// 記事タイトルが同じかどうか比較
	if got.Title != expected.Title {
		t.Errorf("Title: get %s but want %s\n", got.Title, expected.Title)
	}
	// 記事本文が同じかどうか比較
	if got.Contents != expected.Contents {
		t.Errorf("Content: get %s but want %s\n", got.Contents, expected.Contents)
	}
	// 記事投稿者が同じかどうか比較
	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s\n", got.UserName, expected.UserName)
	}
	// 記事いいね数が同じかどうか比較
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, expected.NiceNum)
	}
}

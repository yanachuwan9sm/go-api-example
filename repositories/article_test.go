package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yanachuwan9sm/myapi-tutorial/models"
	"github.com/yanachuwan9sm/myapi-tutorial/repositories"
)

// InsertArticle 関数のテスト
func TestInsertArticle(t *testing.T) {

	// arrange
	article := models.Article{
		Title:    "InsertTest",
		Contents: "testtest",
		UserName: "saki",
	}
	expectedArticleNum := 3

	// act
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}

	// assert
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum,
			newArticle.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
		delete from articles
		where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {

	// テストしたいケースを構造体のスライスの形で作成
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			// 記事 ID1番のテストデータ
			testTitle: "subtest1",
			expected: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  2,
			},
		},

		{
			// 記事 ID2番のテストデータ
			testTitle: "subtest2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum:  4,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {

			// テスト対象となる関数を実行
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				// SelectArticleDetailがうまくいかなくてそもそも返り値gotが得られていないなら
				// この後の期待する出力と実際の出力の比較が不可能なのでテスト続行不可
				t.Fatal(err)
			}

			// 評価

			// 記事IDが同じかどうか比較
			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			// 記事タイトルが同じかどうか比較
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			// 記事本文が同じかどうか比較
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			// 記事投稿者が同じかどうか比較
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			// 記事いいね数が同じかどうか比較
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {

	// arrange
	expectedNum := 2

	// act
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}
	// assert
	if num := len(got); num != expectedNum {
		// SelectArticleList関数から得たArticleスライスの長さが期待通りでないならFAILにする
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

// UpdateNiceNum関数のテスト
func TestUpdateNiceNum(t *testing.T) {

	articleID := 1

	before, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("fail to get before data")
	}

	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("fail to get after data")
	}

	if after.NiceNum-before.NiceNum != 1 {
		t.Error("fail to update nice num")
	}
}

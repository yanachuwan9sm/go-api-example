package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yanachuwan9sm/myapi-tutorial/models"
	"github.com/yanachuwan9sm/myapi-tutorial/repositories"
)

func TestSelectCommentList(t *testing.T) {

	// arrange
	articleID := 1

	// act
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	// assert
	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got ID %d\n", articleID, comment.ArticleID)
		}
	}

}

func TestInsertComment(t *testing.T) {

	// arrange
	comment := models.Comment{
		Message:   "test comment",
		ArticleID: 1,
	}
	expectedCommentNum := 3

	// act
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}

	// assert
	if newComment.CommentID != expectedCommentNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedCommentNum,
			newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
		delete from comments
		where article_id = ? and message = ?
		`
		testDB.Exec(sqlStr, comment.ArticleID, comment.Message)
	})

}

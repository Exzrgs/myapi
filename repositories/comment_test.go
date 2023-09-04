package repositories_test

import (
	"testing"

	"github.com/Exzrgs/myapi/models"
	"github.com/Exzrgs/myapi/repositories"
	"github.com/Exzrgs/myapi/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "testtest",
	}

	expectedID := len(testdata.CommentTestData) + 1

	got, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}

	if got.CommentID != expectedID {
		t.Errorf("CommentID is expected %d but got %d\n", expectedID, got.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
		delete from comments
		where article_id = ? and message = ?;
		`

		testDB.Exec(sqlStr, comment.ArticleID, comment.Message)
	})
}

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	expectedNum := len(testdata.CommentTestData)

	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != expectedNum {
		t.Errorf("CommentNum is expected %d but got %d\n", expectedNum, len(got))
	}
}

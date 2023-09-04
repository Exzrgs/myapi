package repositories_test

import (
	"testing"

	"github.com/Exzrgs/myapi/models"
	"github.com/Exzrgs/myapi/repositories"
	"github.com/Exzrgs/myapi/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetails(t *testing.T) {

	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		// テストを動かす。テストタイトルと、テスト関数を渡して実行する
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID is expected %d but got %d\n", test.expected.ID, got.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("got %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("got %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("got %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("got %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)

	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if expectedNum != len(got) {
		t.Errorf("want %d but got %d articles\n", expectedNum, len(got))
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testtest",
		UserName: "saki",
	}

	expectedNum := len(testdata.ArticleTestData) + 1

	got, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != expectedNum {
		t.Errorf("ArticleID is expected %d but got %d\n", expectedNum, got.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
		delete from articles
		where title = ? and contents = ? and username = ?;
		`

		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestUpdateNiceNum(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		before, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
		if err != nil {
			t.Fatal(err)
		}

		err = repositories.UpdateNiceNum(testDB, test.expected.ID)
		if err != nil {
			t.Fatal(err)
		}

		after, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
		if err != nil {
			t.Fatal(err)
		}

		if after.NiceNum != before.NiceNum+1 {
			t.Errorf("want %d but got %d\n", test.expected.NiceNum+1, after.NiceNum)
		}

		t.Cleanup(func() {
			const sqlStr = `
			update articles
			set nice = ?
			where article_id = ?;
			`

			testDB.Exec(sqlStr, test.expected.NiceNum, test.expected.ID)
		})
	}
}

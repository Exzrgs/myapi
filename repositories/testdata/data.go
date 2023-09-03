package testdata

import (
	"github.com/Exzrgs/myapi/models"
)

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  2,
	},
	models.Article{
		ID:       2,
		Title:    "second",
		Contents: "This is second blog",
		UserName: "saki",
		NiceNum:  1,
	},
}

var CommentTestData = []models.Comment{
	models.Comment{
		ArticleID: 1,
		Message:   "1st comment",
	},
	models.Comment{
		ArticleID: 1,
		Message:   "hello",
	},
}

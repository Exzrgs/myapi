package testdata

import (
	"github.com/Exzrgs/myapi/models"
)

type serviceMoc struct{}

func NewServiceMoc() *serviceMoc {
	return &serviceMoc{}
}

// これらの関数に変な値が放り込まれたときのエラー処理はこの関数で行うので、それはこれらの関数のテストでチェックすればよい。
func (s *serviceMoc) GetArticleService(ID int) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMoc) PostArticleService(reqArticle models.Article) (models.Article, error) {
	return articleTestData[1], nil
}

func (s *serviceMoc) GetArticleListService(page int) ([]models.Article, error) {
	return articleTestData, nil
}

func (s serviceMoc) PostNiceService(article models.Article) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMoc) PostCommentService(reqComment models.Comment) (models.Comment, error) {
	return commentTestData[0], nil
}

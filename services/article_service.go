package services

import (
	"fmt"

	"github.com/Exzrgs/myapi/models"
	"github.com/Exzrgs/myapi/repositories"
)

func (s *MyAppService) GetArticleService(ID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, ID)
	if err != nil {
		fmt.Println("error at SelectArticleDetail in GetArticleService")
		return models.Article{}, err
	}

	comments, err := repositories.SelectCommentList(s.db, ID)
	if err != nil {
		fmt.Println("error at SelectCommentList in GetArticleService")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, comments...)

	return article, nil
}

func (s *MyAppService) PostArticleService(reqArticle models.Article) (models.Article, error) {
	resArticle, err := repositories.InsertArticle(s.db, reqArticle)
	if err != nil {
		fmt.Println("error at InsertArticle in PostArticleService")
		return models.Article{}, err
	}

	return resArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		fmt.Println("error at SelectArticleList in GetArticleListService")
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	// デバッグ用
	// fmt.Printf("reqArticle is %+v in PostNiceService\n", article)

	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		fmt.Println("error at UpdateNiceNum in PostNiceService")
		return models.Article{}, err
	}

	newArticle, err := repositories.SelectArticleDetail(s.db, article.ID)
	if err != nil {
		fmt.Println("error at SelectArticleDetail in PostNiceService")
		return models.Article{}, err
	}

	return newArticle, nil
}

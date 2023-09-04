package services

import (
	"fmt"

	"github.com/Exzrgs/myapi/models"
	"github.com/Exzrgs/myapi/repositories"
)

func GetArticleService(ID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		fmt.Println("error at connectDB in GetArticleService")
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, ID)
	if err != nil {
		fmt.Println("error at SelectArticleDetail in GetArticleService")
		return models.Article{}, err
	}

	comments, err := repositories.SelectCommentList(db, ID)
	if err != nil {
		fmt.Println("error at SelectCommentList in GetArticleService")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, comments...)

	return article, nil
}

func PostArticleService(reqArticle models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		fmt.Println("error at connectDB in PostArticleService")
		return models.Article{}, err
	}

	resArticle, err := repositories.InsertArticle(db, reqArticle)
	if err != nil {
		fmt.Println("error at InsertArticle in PostArticleService")
		return models.Article{}, err
	}

	return resArticle, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		fmt.Println("error at connectDB in GetArticleListService")
		return nil, err
	}

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		fmt.Println("error at SelectArticleList in GetArticleListService")
		return nil, err
	}

	return articleList, nil
}

func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		fmt.Println("error at connectDB in PostNiceService")
		return models.Article{}, err
	}

	// デバッグ用
	// fmt.Printf("reqArticle is %+v in PostNiceService\n", article)

	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		fmt.Println("error at UpdateNiceNum in PostNiceService")
		return models.Article{}, err
	}

	newArticle, err := repositories.SelectArticleDetail(db, article.ID)
	if err != nil {
		fmt.Println("error at SelectArticleDetail in PostNiceService")
		return models.Article{}, err
	}

	return newArticle, nil
}

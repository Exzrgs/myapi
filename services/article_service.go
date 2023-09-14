package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Exzrgs/myapi/apperrors"
	"github.com/Exzrgs/myapi/models"
	"github.com/Exzrgs/myapi/repositories"
)

func (s *MyAppService) GetArticleService(ID int) (models.Article, error) {
	var article models.Article
	var comments []models.Comment
	var aErr error
	var cErr error

	type gotArticle struct {
		article models.Article
		err     error
	}
	aCh := make(chan gotArticle)
	defer close(aCh)

	type gotComments struct {
		comments []models.Comment
		err      error
	}
	cCh := make(chan gotComments)
	defer close(cCh)

	go func(ch chan<- gotArticle, db *sql.DB, article models.Article) {
		article, err := repositories.SelectArticleDetail(db, ID)
		resArticle := gotArticle{article: article, err: err}
		ch <- resArticle
	}(aCh, s.db, article)

	go func(ch chan<- gotComments, db *sql.DB, cooments []models.Comment) {
		comments, err := repositories.SelectCommentList(db, ID)
		resComments := gotComments{comments: comments, err: err}
		ch <- resComments
	}(cCh, s.db, comments)

	for i := 0; i < 2; i++ {
		select {
		case gotArt := <-aCh:
			article, aErr = gotArt.article, gotArt.err
		case gotCom := <-cCh:
			comments, cErr = gotCom.comments, gotCom.err
		}
	}

	if aErr != nil {
		fmt.Println("error at SelectArticleDetail in GetArticleService")

		if errors.Is(aErr, sql.ErrNoRows) {
			aErr = apperrors.NAData.Wrap(aErr, "no data")
		} else {
			aErr = apperrors.GetDataFailed.Wrap(aErr, "fail to get data")
		}

		return models.Article{}, aErr
	}

	if cErr != nil {
		fmt.Println("error at SelectCommentList in GetArticleService")
		cErr = apperrors.GetDataFailed.Wrap(cErr, "fail to get data")
		return models.Article{}, cErr
	}

	article.CommentList = append(article.CommentList, comments...)

	return article, nil
}

func (s *MyAppService) PostArticleService(reqArticle models.Article) (models.Article, error) {
	resArticle, err := repositories.InsertArticle(s.db, reqArticle)
	if err != nil {
		fmt.Println("error at InsertArticle in PostArticleService")
		err = apperrors.InsertDataFailed.Wrap(err, "fail to recode data")
		return models.Article{}, err
	}

	return resArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		fmt.Println("error at SelectArticleList in GetArticleListService")
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err = apperrors.NAData.Wrap(ErrorNoData, "no data")
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

		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "no data")
		} else {
			err = apperrors.UpdateDataFailed.Wrap(err, "fail to update data")
		}

		return models.Article{}, err
	}

	newArticle, err := repositories.SelectArticleDetail(s.db, article.ID)
	if err != nil {
		fmt.Println("error at SelectArticleDetail in PostNiceService")
		return models.Article{}, err
	}

	return newArticle, nil
}

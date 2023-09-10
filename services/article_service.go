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
	article, err := repositories.SelectArticleDetail(s.db, ID)
	if err != nil {
		fmt.Println("error at SelectArticleDetail in GetArticleService")

		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
		} else {
			err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		}

		return models.Article{}, err
	}

	comments, err := repositories.SelectCommentList(s.db, ID)
	if err != nil {
		fmt.Println("error at SelectCommentList in GetArticleService")
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
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

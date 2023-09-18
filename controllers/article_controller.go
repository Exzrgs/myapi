package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Exzrgs/myapi/apperrors"
	"github.com/Exzrgs/myapi/common"
	"github.com/Exzrgs/myapi/controllers/services"
	"github.com/Exzrgs/myapi/models"
	"github.com/gorilla/mux"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

// ブログ記事の投稿をする
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	authedUserName := common.GetUserName(req.Context())
	if authedUserName != reqArticle.UserName {
		err := apperrors.NotMatchUser.Wrap(errors.New("not match authed user and post user"), "invalid parameter")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// ブログ記事の一覧を取得
func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int

	p, ok := queryMap["page"]

	if ok && len(p) > 0 {
		var err error

		page, err = strconv.Atoi(p[0])

		if err != nil {
			err = apperrors.BadParameter.Wrap(err, "query parameter must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

// 記事ナンバーID番の登校データを取得
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParameter.Wrap(err, "query parameter must be number")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// 記事にいいねをつける
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}
	// デバッグ用
	// fmt.Printf("reqArticle is %+v in PostNiceHandler\n", reqArticle)

	resArticle, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(resArticle)
}

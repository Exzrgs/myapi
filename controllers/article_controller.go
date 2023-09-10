package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
		http.Error(w, "fail to decode json\n", http.StatusInternalServerError)
	}

	resArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resArticle); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}
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
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(articleList); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}
}

// 記事ナンバーID番の登校データを取得
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}
}

// 記事にいいねをつける
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusInternalServerError)
	}
	// デバッグ用
	// fmt.Printf("reqArticle is %+v in PostNiceHandler\n", reqArticle)

	resArticle, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		fmt.Println("error at PostNiceService in PostNiceHandler")
		fmt.Println(err)
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resArticle); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}
}

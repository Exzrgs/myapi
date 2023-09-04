package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Exzrgs/myapi/models"
	"github.com/Exzrgs/myapi/services"
	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello\n")
}

// ブログ記事の投稿をする
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusInternalServerError)
	}

	resArticle, err := services.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resArticle); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}
}

// ブログ記事の一覧を取得
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
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

	articleList, err := services.GetArticleListService(page)
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
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := services.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}
}

// 記事にいいねをつける
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusInternalServerError)
	}
	// デバッグ用
	// fmt.Printf("reqArticle is %+v in PostNiceHandler\n", reqArticle)

	resArticle, err := services.PostNiceService(reqArticle)
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

// 記事にコメントをつける
func CommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusInternalServerError)
	}

	resComment, err := services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resComment); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}
}

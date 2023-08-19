package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Exzrgs/myapi/models"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello\n")
}

// ブログ記事の投稿をする
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "Posting Article...\n")

	// var reqBodybuffer []byte

	length, err := strconv.Atoi(req.Header.Get("Content-Length"))

	if err != nil {
		http.Error(w, "cannot get content length\n", http.StatusBadRequest)
		return
	}

	reqBodybuffer := make([]byte, length)

	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	///////////////////////////////

	article := models.Article1
	jsonData, err := json.Marshal(article)

	if err != nil {
		http.Error(w, "fail to encode to json\n", http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

// ブログ記事の一覧を取得
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	/*
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

		resString := fmt.Sprintf("Article List (page %d)\n", page)
		io.WriteString(w, resString)
	*/
	ls := []models.Article{}
	ls = append(ls, models.Article1, models.Article2)

	jsonData, err := json.Marshal(ls)

	if err != nil {
		http.Error(w, "fail to encode to json\n", http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

// 記事ナンバーID番の登校データを取得
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	/*
		articleID, err := strconv.Atoi(mux.Vars(req)["id"])

		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}

		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	*/

	article := models.Article1

	jsonData, err := json.Marshal(article)

	if err != nil {
		http.Error(w, "fail to encode to json\n", http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

// 記事にいいねをつける
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "Posting Nice...\n")

	article := models.Article1

	jsonData, err := json.Marshal(article)

	if err != nil {
		http.Error(w, "fail to encode to json\n", http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

// 記事にコメントをつける
func CommentHandler(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "Posting Comment...\n")

	comment := models.Comment1

	jsonData, err := json.Marshal(comment)

	if err != nil {
		http.Error(w, "fail to encode to json\n", http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

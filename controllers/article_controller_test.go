package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// 引数は*testing.T
func TestArticleListHandler(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		resCode int
	}{
		{
			name:  "number",
			query: "1",
			// 数字じゃなくて、メソッドを使う
			resCode: http.StatusOK,
		},
		{
			name:    "string",
			query:   "aaa",
			resCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// URLは一度変数にしたほうがいい
			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s", test.query)
			// POSTとかだったらボディが必要だけど、GETならいらない
			req := httptest.NewRequest("GET", url, nil)

			res := httptest.NewRecorder()

			// handlerはサブルーチンではなくメソッド
			aCon.ArticleListHandler(res, req)

			if res.Code != test.resCode {
				// fmtではなくt.Errorf
				t.Errorf("status code is expected %d but got %d", test.resCode, res.Code)
				return
			}
		})
	}
}

/*
テストの目的はハンドラがうまく動作しているかを見ること。
今回は、ハンドラを明示的に呼び出しはしないが、ルータを設定し、HTTPリクエストを送ることで、そのレスポンスの内容を検証する。
*/
func TestArtcleDetailHandler(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		resCode int
	}{
		{
			name:    "number",
			query:   "1",
			resCode: http.StatusOK,
		},
		{
			name:    "string",
			query:   "aaa",
			resCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := httptest.NewRecorder()

			url := fmt.Sprintf("http://localhost:8080/article/%s", test.query)
			req := httptest.NewRequest("GET", url, nil)

			r := mux.NewRouter()
			r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
			r.ServeHTTP(res, req)

			if res.Code != test.resCode {
				t.Errorf("status code is expected %d but got %d\n", test.resCode, res.Code)
			}
		})
	}
}

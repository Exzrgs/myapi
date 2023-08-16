package main

import (
	"log"
	"net/http"

	"github.com/Exzrgs/myapi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/ID", handlers.ArticleDetailHandler)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.CommentHandler)

	log.Println("running")
	log.Fatal(http.ListenAndServe(":8080", r))
}

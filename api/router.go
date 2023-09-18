package api

import (
	"database/sql"
	"net/http"

	"github.com/Exzrgs/myapi/api/middlewares"
	"github.com/Exzrgs/myapi/controllers"
	"github.com/Exzrgs/myapi/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)

	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.CommentHandler).Methods(http.MethodPost)

	// 順番が大事
	r.Use(middlewares.LoggingMiddleWare)
	r.Use(middlewares.AuthMiddleware)

	return r
}

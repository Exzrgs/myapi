package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Exzrgs/myapi/controllers/services"
	"github.com/Exzrgs/myapi/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// 記事にコメントをつける
func (c *CommentController) CommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusInternalServerError)
	}

	resComment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resComment); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}
}

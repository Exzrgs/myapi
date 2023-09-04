package services

import (
	"fmt"

	"github.com/Exzrgs/myapi/models"
	"github.com/Exzrgs/myapi/repositories"
)

func PostCommentService(reqComment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		fmt.Println("error at connectDB in PostCommentService")
		return models.Comment{}, err
	}

	resComment, err := repositories.InsertComment(db, reqComment)
	if err != nil {
		fmt.Println("error at InsertComment in PostCommentService")
		return models.Comment{}, err
	}

	return resComment, nil
}

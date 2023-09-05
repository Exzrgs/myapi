package services

import (
	"fmt"

	"github.com/Exzrgs/myapi/models"
	"github.com/Exzrgs/myapi/repositories"
)

func (s *MyAppService) PostCommentService(reqComment models.Comment) (models.Comment, error) {
	resComment, err := repositories.InsertComment(s.db, reqComment)
	if err != nil {
		fmt.Println("error at InsertComment in PostCommentService")
		return models.Comment{}, err
	}

	return resComment, nil
}

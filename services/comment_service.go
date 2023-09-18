package services

import (
	"github.com/Exzrgs/myapi/apperrors"
	"github.com/Exzrgs/myapi/models"
	"github.com/Exzrgs/myapi/repositories"
)

func (s *MyAppService) PostCommentService(reqComment models.Comment) (models.Comment, error) {
	resComment, err := repositories.InsertComment(s.db, reqComment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return resComment, nil
}

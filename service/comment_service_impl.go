package service

import (
	"context"
	"database/sql"

	"github.com/MeizalunaWulandari/go-artkl/helper"
	"github.com/MeizalunaWulandari/go-artkl/model/domain"
	"github.com/MeizalunaWulandari/go-artkl/model/web"
	"github.com/MeizalunaWulandari/go-artkl/repository"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	DB                *sql.DB
}

func (service *CommentServiceImpl) Create(ctx context.Context, request web.CommentCreateRequest) web.CommentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	comment := domain.Comment{
		Content:   request.Content,
		UserId:    request.UserId,
		ArticleId: request.ArticleId,
	}

	comment = service.CommentRepository.Save(ctx, tx, comment)
	return helper.ToCommentResponse(comment)
}
func (service *CommentServiceImpl) Update(ctx context.Context, request web.CommentUpdateRequest) web.CommentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	comment, err := service.CommentRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	comment.Content = request.Content
	comment.UserId = request.UserId
	comment.ArticleId = request.ArticleId

	comment = service.CommentRepository.Update(ctx, tx, comment)
	return helper.ToCommentResponse(comment)

}

func (service *CommentServiceImpl) Delete(ctx context.Context, commentId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	comment, err := service.CommentRepository.FindById(ctx, tx, commentId)
	helper.PanicIfError(err)

	service.CommentRepository.Delete(ctx, tx, comment)

}
func (service *CommentServiceImpl) FindById(ctx context.Context, commentId int) web.CommentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	comment, err := service.CommentRepository.FindById(ctx, tx, commentId)
	helper.PanicIfError(err)
	return helper.ToCommentResponse(comment)
}

func (service *CommentServiceImpl) FindAll(ctx context.Context) []web.CommentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	comments := service.CommentRepository.FindAll(ctx, tx)

	return helper.ToCommentResponses(comments)

}

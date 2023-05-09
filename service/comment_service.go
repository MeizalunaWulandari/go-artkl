package service

import (
	"context"

	"github.com/MeizalunaWulandari/go-artkl/model/web"
)

type CommentService interface {
	Create(ctx context.Context, request web.CommentCreateRequest) web.CommentResponse
	Update(ctx context.Context, request web.ArticleUpdateRequest) web.CommentResponse
	Delete(ctx context.Context, commentId int)
	FindById(ctx context.Context, commentId int) web.CommentResponse
	FindAll(ctx context.Context) []web.CommentResponse
}

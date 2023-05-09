package service

import (
	"context"

	"github.com/MeizalunaWulandari/go-artkl/model/web"
)

type ArticleService interface {
	Create(ctx context.Context, request web.ArticleCreateRequest) web.ArticleResponse
	Update(ctx context.Context, request web.ArticleUpdateRequest) web.ArticleResponse
	Delete(ctx context.Context, articleId int)
	FindById(ctx context.Context, articleIdId int) web.ArticleResponse
	FindAll(ctx context.Context) []web.ArticleResponse
}

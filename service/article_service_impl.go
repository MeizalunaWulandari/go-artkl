package service

import (
	"context"
	"database/sql"

	"github.com/MeizalunaWulandari/go-artkl/helper"
	"github.com/MeizalunaWulandari/go-artkl/model/domain"
	"github.com/MeizalunaWulandari/go-artkl/model/web"
	"github.com/MeizalunaWulandari/go-artkl/repository"
)

type ArticleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
	DB                *sql.DB
}

func (service *ArticleServiceImpl) Create(ctx context.Context, request web.ArticleCreateRequest) web.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	article := domain.Article{
		Slug:       request.Slug,
		Title:      request.Title,
		Content:    request.Slug,
		CategoryId: request.CategoryId,
		UserId:     request.UserId,
	}
	article = service.ArticleRepository.Save(ctx, tx, article)
	return helper.ToArticleResponse(article)
}

func (service *ArticleServiceImpl) Update(ctx context.Context, request web.ArticleUpdateRequest) web.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.ArticleRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	article.Slug = request.Slug
	article.Title = request.Title
	article.Content = request.Content
	article.CategoryId = request.CategoryId
	article.UserId = request.UserId

	article = service.ArticleRepository.Update(ctx, tx, article)
	return helper.ToArticleResponse(article)

}

// Harusnya 2 parameter
func (service *ArticleServiceImpl) Delete(ctx context.Context, articleId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.ArticleRepository.FindById(ctx, tx, articleId)
	helper.PanicIfError(err)

	service.ArticleRepository.Delete(ctx, tx, article)

}

func (service *ArticleServiceImpl) FindById(ctx context.Context, articleId int) web.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.ArticleRepository.FindById(ctx, tx, articleId)
	helper.PanicIfError(err)
	return helper.ToArticleResponse(article)
}

func (service *ArticleServiceImpl) FindAll(ctx context.Context) []web.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	articles := service.ArticleRepository.FindAll(ctx, tx)

	return helper.ToArticleResponses(articles)
}

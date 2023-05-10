package helper

import (
	"github.com/MeizalunaWulandari/go-artkl/model/domain"
	"github.com/MeizalunaWulandari/go-artkl/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToArticleResponse(article domain.Article) web.ArticleResponse {
	return web.ArticleResponse{
		Id:         article.Id,
		Slug:       article.Slug,
		Title:      article.Title,
		Content:    article.Content,
		CategoryId: article.CategoryId,
		UserId:     article.UserId,
		Status:     article.Status,
		Views:      article.Views,
	}
}

func ToArticleResponses(articles []domain.Article) []web.ArticleResponse {
	var articleResponses []web.ArticleResponse
	for _, article := range articles {
		articleResponses = append(articleResponses, ToArticleResponse(article))
	}
	return articleResponses
}

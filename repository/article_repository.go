package repository

import (
	"context"
	"database/sql"

	"github.com/MeizalunaWulandari/go-artkl/model/domain"
)

type ArticleRepository interface {
	Save(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article
	Update(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article
	Delete(ctx context.Context, tx *sql.Tx, article domain.Article)
	FindBySlug(ctx context.Context, tx *sql.Tx, articleSlug string) domain.Article
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Article
	SortByPopularity(ctx context.Context, tx *sql.Tx) []domain.Article
	SortByWriter(ctx context.Context, tx *sql.Tx, articleWriter string) []domain.Article
	SortByCategory(ctx context.Context, tx *sql.Tx, articleCategory string) []domain.Article
}

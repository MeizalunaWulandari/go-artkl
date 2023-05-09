package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MeizalunaWulandari/go-artkl/helper"
	"github.com/MeizalunaWulandari/go-artkl/model/domain"
)

type ArticleRepositoryImpl struct {
}

func (repository *ArticleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article {
	SQL := "INSERT INTO article(slug, title, content, category_id, user_id) VALUES(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, article.Slug, article.Title, article.CategoryId, article.UserId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	article.Id = int(id)
	return article
}
func (repository *ArticleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article {
	SQL := "UPDATE article SET slug = ?, title = ?, content = ?, category_id = ? WHERE id = ? AND user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, article.Slug, article.Title, article.CategoryId, article.Id, article.UserId)
	helper.PanicIfError(err)
	return article
}
func (repository *ArticleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, article domain.Article) {
	SQL := "DELETE FROM article WHERE id = ? AND user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, article.Id, article.UserId)
	helper.PanicIfError(err)
}
func (repository *ArticleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, articleId int) (domain.Article, error) {
	SQL := "SELECT, id, slug, title, content, category_id, user_id, views, created_at FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, articleId)
	helper.PanicIfError(err)
	article := domain.Article{}
	if rows.Next() {
		err := rows.Scan(&article.Id, &article.Slug, &article.Title, article.Content, article.CategoryId, article.UserId, article.Views, article.CreatedAt)
		helper.PanicIfError(err)
		return article, nil
	} else {
		return article, errors.New("category is not found")
	}
}

func (repository *ArticleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, article domain.Article) []domain.Article {
	SQL := "SELECT, id, slug, title, content, category_id, user_id, views, created_at FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var articles []domain.Article
	for rows.Next() {
		article := domain.Article{}
		err := rows.Scan(&article.Id, &article.Slug, &article.Title, &article.Content, &article.CategoryId, &article.UserId, article.Views, article.CreatedAt)
		helper.PanicIfError(err)
		articles = append(articles, article)
	}
	return articles
}

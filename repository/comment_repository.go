package repository

import (
	"context"
	"database/sql"

	"github.com/MeizalunaWulandari/go-artkl/model/domain"
)

type CommentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment
	Update(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment
	Delete(ctx context.Context, tx *sql.Tx, comment domain.Comment)
	FindById(ctx context.Context, tx *sql.Tx, commentId int) (domain.Comment, error)
	FindByArticle(ctx context.Context, tx *sql.Tx, articleId int) domain.Comment
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Comment
}

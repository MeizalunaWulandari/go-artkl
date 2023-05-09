package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MeizalunaWulandari/go-artkl/helper"
	"github.com/MeizalunaWulandari/go-artkl/model/domain"
)

type CommentRepositoryImpl struct {
}

func (repository *CommentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment {
	SQL := "INSERT INTO comments(content, user_id, article_id) VALUES(?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, comment.Content, comment.UserId, comment.ArticleId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	comment.Id = int(id)
	return comment
}
func (repository *CommentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment {
	SQL := "UPDATE comments SET content = ? WHERE id = ? AND user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, comment.Content, comment.Id, comment.UserId)
	helper.PanicIfError(err)
	return comment
}
func (repository *CommentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, comment domain.Comment) {
	SQL := "DELETE FROM comments WHERE id = ? AND user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, comment.Id, comment.UserId)
	helper.PanicIfError(err)
}
func (repository *CommentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, commentId int) (domain.Comment, error) {
	SQL := "SELECT, id, content, user_id, article_id, created_at FROM comments WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, commentId)
	helper.PanicIfError(err)
	comment := domain.Comment{}
	if rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Content, &comment.UserId, &comment.ArticleId, comment.CreatedAt)
		helper.PanicIfError(err)
		return comment, nil
	} else {
		return comment, errors.New("comment is not found")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, comment domain.Comment) []domain.Comment {

	SQL := "SELECT, id, content, user_id, article_id, created_at FROM comments"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var comments []domain.Comment
	for rows.Next() {
		comment := domain.Comment{}
		err := rows.Scan(&comment.Id, &comment.Content, &comment.UserId, &comment.ArticleId, &comment.CreatedAt)
		helper.PanicIfError(err)
		comments = append(comments, comment)
	}
	return comments
}

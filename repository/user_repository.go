package repository

import (
	"context"
	"database/sql"

	"github.com/MeizalunaWulandari/go-artkl/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) domain.User
	FindAll(ctx context.Context, tx *sql.Tx, user domain.User) []domain.User
}

package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MeizalunaWulandari/go-artkl/helper"
	"github.com/MeizalunaWulandari/go-artkl/model/domain"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(name, email, username, password) VALUES(?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Username, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}
func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET name = ?, email = ?, username = ?, password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Username, user.Password, user.Id)
	helper.PanicIfError(err)
	return user
}
func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}
func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT, id, name, email, username, role, level FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Role, &user.Level)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, user domain.User) []domain.User {
	SQL := "SELECT, id, name, email, username, role, level, created_at FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Username, &user.Role, &user.Level, &user.CreateAt)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}

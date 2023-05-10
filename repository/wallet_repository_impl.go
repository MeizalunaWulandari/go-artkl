package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MeizalunaWulandari/go-artkl/helper"
	"github.com/MeizalunaWulandari/go-artkl/model/domain"
)

type WalletRepositoryImpl struct {
}

func (repository *WalletRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) domain.Wallet {
	SQL := "INSERT INTO wallets(balance, user_id, status) VALUES(?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, wallet.Balance, wallet.UserId, wallet.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	wallet.Id = int(id)
	return wallet
}
func (repository *WalletRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) domain.Wallet {
	SQL := "UPDATE wallets SET balance = ?, status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, wallet.Balance, wallet.Status, wallet.Id)
	helper.PanicIfError(err)
	return wallet
}
func (repository *WalletRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) {
	SQL := "DELETE FROM wallets WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, wallet.Id)
	helper.PanicIfError(err)
}
func (repository *WalletRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, walletId int) (domain.Wallet, error) {
	SQL := "SELECT, id, balance, user_id, status name FROM wallets WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, walletId)
	helper.PanicIfError(err)
	wallet := domain.Wallet{}
	if rows.Next() {
		err := rows.Scan(&wallet.Id, &wallet.Balance, &wallet.UserId, wallet.Status, wallet.Id)
		helper.PanicIfError(err)
		return wallet, nil
	} else {
		return wallet, errors.New("category is not found")
	}
}

func (repository *WalletRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Wallet {
	SQL := "SELECT, id, balance, user_id, status name FROM wallets"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var wallets []domain.Wallet
	for rows.Next() {
		wallet := domain.Wallet{}
		err := rows.Scan(&wallet.Id, &wallet.Balance, &wallet.UserId, wallet.Status, wallet.Id)
		helper.PanicIfError(err)
		wallets = append(wallets, wallet)
	}
	return wallets
}

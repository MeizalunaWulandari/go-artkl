package repository

import (
	"context"
	"database/sql"

	"github.com/MeizalunaWulandari/go-artkl/model/domain"
)

type WalletRepository interface {
	Save(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) domain.Wallet
	Update(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) domain.Wallet
	Delete(ctx context.Context, tx *sql.Tx, wallet domain.Wallet)
	FindById(ctx context.Context, tx *sql.Tx, walletId int) (domain.Wallet, error)
	FindAll(ctx context.Context, tx *sql.Tx, wallet domain.Wallet) []domain.Wallet
}

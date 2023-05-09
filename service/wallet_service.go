package service

import (
	"context"

	"github.com/MeizalunaWulandari/go-artkl/model/web"
)

type WalletService interface {
	Create(ctx context.Context, request web.WalletCreateRequest) web.WalletResponse
	Update(ctx context.Context, request web.WalletUpdateRequest) web.WalletResponse
	Delete(ctx context.Context, walletId int)
	FindById(ctx context.Context, walletId int) web.WalletResponse
	FindAll(ctx context.Context) []web.WalletResponse
}

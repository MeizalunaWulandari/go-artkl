package service

import (
	"context"
	"database/sql"

	"github.com/MeizalunaWulandari/go-artkl/helper"
	"github.com/MeizalunaWulandari/go-artkl/model/domain"
	"github.com/MeizalunaWulandari/go-artkl/model/web"
	"github.com/MeizalunaWulandari/go-artkl/repository"
)

type WalletServiceImpl struct {
	WalletRepository repository.WalletRepository
	DB               *sql.DB
}

func (service *WalletServiceImpl) Create(ctx context.Context, request web.WalletCreateRequest) web.WalletResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	wallet := domain.Wallet{
		Balance: request.Balance,
		UserId:  request.UserId,
	}
	wallet = service.WalletRepository.Save(ctx, tx, wallet)
	return helper.ToWalletResponse(wallet)
}

func (service *WalletServiceImpl) Update(ctx context.Context, request web.WalletUpdateRequest) web.WalletResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	wallet, err := service.WalletRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	wallet.Balance = request.Balance
	wallet.Status = request.Status

	wallet = service.WalletRepository.Update(ctx, tx, wallet)
	return helper.ToWalletResponse(wallet)

}

// Harusnya 2 parameter
func (service *WalletServiceImpl) Delete(ctx context.Context, walletId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	wallet, err := service.WalletRepository.FindById(ctx, tx, walletId)
	helper.PanicIfError(err)

	service.WalletRepository.Delete(ctx, tx, wallet)

}

func (service *WalletServiceImpl) FindById(ctx context.Context, walletId int) web.WalletResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	wallet, err := service.WalletRepository.FindById(ctx, tx, walletId)
	helper.PanicIfError(err)
	return helper.ToWalletResponse(wallet)
}

func (service *WalletServiceImpl) FindAll(ctx context.Context) []web.WalletResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	wallets := service.WalletRepository.FindAll(ctx, tx)

	return helper.ToWalletResponses(wallets)
}

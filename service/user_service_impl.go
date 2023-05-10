package service

import (
	"context"
	"database/sql"

	"github.com/MeizalunaWulandari/go-artkl/helper"
	"github.com/MeizalunaWulandari/go-artkl/model/domain"
	"github.com/MeizalunaWulandari/go-artkl/model/web"
	"github.com/MeizalunaWulandari/go-artkl/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	}
	user = service.UserRepository.Save(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.Name = request.Name
	user.Email = request.Email
	user.Username = request.Username
	user.Password = request.Password

	user = service.UserRepository.Update(ctx, tx, user)
	return helper.ToUserResponse(user)

}

// Harusnya 2 parameter
func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	service.UserRepository.Delete(ctx, tx, article)

}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(user)
}

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	forgotPasswordRepo "golang-bootcamp-1/internal/forgot-password/repository"
	forgotPasswordH "golang-bootcamp-1/internal/forgot-password/transport/http"
	forgotPasswordUc "golang-bootcamp-1/internal/forgot-password/usecase"
	oauthRepo "golang-bootcamp-1/internal/oauth/repository"
	oauthH "golang-bootcamp-1/internal/oauth/transport/http"
	oauthUC "golang-bootcamp-1/internal/oauth/usecase"
	registerH "golang-bootcamp-1/internal/register/transport/http"
	registerUc "golang-bootcamp-1/internal/register/usecase"
	userUC "golang-bootcamp-1/internal/user/usecase"
	mailUc "golang-bootcamp-1/pkg/mail/mailtrap"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// func InitializeUserUsecase(db *gorm.DB) userUC.IUserUseCase {
// 	wire.Build(userUC.NewUserUseCase, userRepo.NewUserRepo)
// 	return userUC.IUserUseCase
// }

func InitializeRegisterHandler(db *gorm.DB, userUc userUC.IUserUseCase) *registerH.RegisterHandler {
	wire.Build(registerH.NewRegisterHandler, registerUc.NewRegisterUseCase, mailUc.NewMailUsecase)
	return &registerH.RegisterHandler{}
}

func InitializeOauthHandler(db *gorm.DB, userUc userUC.IUserUseCase) *oauthH.OauthHandler {
	wire.Build(oauthH.NewOauthHandler, oauthUC.NewOauthUseCase, oauthRepo.NewOauthRefreshTokenRepo, oauthRepo.NewOauthAcces, oauthRepo.NewOauthClientRepo)
	return &oauthH.OauthHandler{}
}

func InitializeForgotPasswordHanlder(db *gorm.DB, userUc userUC.IUserUseCase) *forgotPasswordH.ForgotPasswordHandler {
	wire.Build(
		forgotPasswordH.NewForgotPasswordHandler,
		forgotPasswordUc.NewForgotPasswordUsecase,
		mailUc.NewMailUsecase,
		forgotPasswordRepo.NewForgotPasswordRepository,
	)

	return &forgotPasswordH.ForgotPasswordHandler{}
}

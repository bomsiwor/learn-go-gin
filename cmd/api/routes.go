package main

import (
	userRepo "golang-bootcamp-1/internal/user/repository"
	userUC "golang-bootcamp-1/internal/user/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initHandler(db *gorm.DB, router *gin.RouterGroup) {
	userRepo := userRepo.NewUserRepo(db)
	userUc := userUC.NewUserUseCase(userRepo)

	InitializeRegisterHandler(db, userUc).Router(router)
	InitializeOauthHandler(db, userUc).Router(router)
	InitializeForgotPasswordHanlder(db, userUc).Router(router)
}

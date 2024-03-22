package main

import (
	"golang-bootcamp-1/config"
	oauthRepo "golang-bootcamp-1/internal/oauth/repository"
	oauthH "golang-bootcamp-1/internal/oauth/transport/http"
	oauthUC "golang-bootcamp-1/internal/oauth/usecase"
	registerH "golang-bootcamp-1/internal/register/transport/http"
	registerUc "golang-bootcamp-1/internal/register/usecase"
	userRepo "golang-bootcamp-1/internal/user/repository"
	userUC "golang-bootcamp-1/internal/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	// Load env at main
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	db := config.InitDB()

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "Hello world!")
		})
	}

	userRepo := userRepo.NewUserRepo(db)
	userUC := userUC.NewUserUseCase(userRepo)

	registerUc := registerUc.NewRegisterUseCase(userUC)
	registerHandler := registerH.NewRegisterHandler(registerUc)
	registerHandler.Router(&r.RouterGroup)

	// Login
	oauthClientRepo := oauthRepo.NewOauthClientRepo(db)
	oauthAccessToken := oauthRepo.NewOauthAcces(db)
	oauthRefreshToken := oauthRepo.NewOauthRefreshTokenRepo(db)
	oauthUsecase := oauthUC.NewOauthUseCase(
		oauthClientRepo,
		oauthAccessToken,
		oauthRefreshToken,
		userUC,
	)
	oauthHandler := oauthH.NewOauthHandler(oauthUsecase)
	oauthHandler.Router(&r.RouterGroup)

	r.Run("127.0.0.1:8082")
}

package main

import (
	"golang-bootcamp-1/config"
	userRepo "golang-bootcamp-1/internal/user/repository"
	userUC "golang-bootcamp-1/internal/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	// _ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	// Load env at main
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// environtmentPath := filepath.Join(dir, ".env")
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
	initHandler(db, v1)

	r.Run("127.0.0.1:8082")
}

func initHandler(db *gorm.DB, router *gin.RouterGroup) {
	userRepo := userRepo.NewUserRepo(db)
	userUc := userUC.NewUserUseCase(userRepo)

	InitializeRegisterHandler(db, userUc).Router(router)
	InitializeOauthHandler(db, userUc).Router(router)
}

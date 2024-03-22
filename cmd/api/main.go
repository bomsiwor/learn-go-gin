package main

import (
	"golang-bootcamp-1/config"
	userRepo "golang-bootcamp-1/internal/user/repository"
	userUC "golang-bootcamp-1/internal/user/usecase"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// Load env at main
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	environtmentPath := filepath.Join(dir, ".env")
	err = godotenv.Load(environtmentPath)
	if err != nil {
		panic(err.Error())
	}

	db := config.InitDB()
	initHandler(db, r)

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "Hello world!")
		})
	}

	r.Run("127.0.0.1:8082")
}

func initHandler(db *gorm.DB, engine *gin.Engine) {
	userRepo := userRepo.NewUserRepo(db)
	userUC := userUC.NewUserUseCase(userRepo)

	// Register
	registerHn := InitializeRegisterHandler(db, userUC)
	registerHn.Router(&engine.RouterGroup)

	// Login
	oauthHn := InitializeOauthHandler(db, userUC)
	oauthHn.Router(&engine.RouterGroup)
}

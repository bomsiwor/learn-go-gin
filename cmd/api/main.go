package main

import (
	"golang-bootcamp-1/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// _ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	// CORS config
	cors := cors.Default()
	r.Use(cors)

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
	// Ini handler from routes
	initHandler(db, v1)

	r.Run("127.0.0.1:8082")
}

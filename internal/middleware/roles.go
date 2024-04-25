package middleware

import (
	"golang-bootcamp-1/pkg/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Roles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(roles)
		log.Println(c.Get("user"))
		c.Next()
	}
}

func Permission(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(permissions)
		log.Println(c.Get("user"))
		c.Next()
	}
}

func IsAdmin(c *gin.Context) {
	if !c.GetBool("isAdmin") {
		code := http.StatusUnauthorized

		c.JSON(
			code,
			response.GenerateResponse(
				code,
				http.StatusText(code),
				"Unauthorized",
			),
		)
		c.Abort()
		return
	}

	c.Next()
}

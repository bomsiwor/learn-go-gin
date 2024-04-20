package middleware

import (
	"errors"
	"golang-bootcamp-1/pkg/jwt"
	"golang-bootcamp-1/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware for JWT based authentication,
// protecting route from being accessed without JWT Authorization token
func JwtTokenCheck(c *gin.Context) {
	token, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			response.GenerateResponse(
				http.StatusUnauthorized,
				"Failed to authenticate",
				nil,
			),
		)
		c.Abort()
		return
	}

	claim, err := jwt.ValidateToken(token)
	if err != nil {
		c.JSON(
			http.StatusForbidden,
			response.GenerateResponse(
				http.StatusForbidden,
				"Failed to authenticate",
				nil,
			),
		)
		c.Abort()
		return
	}

	// Set user id to context
	c.Set("user", claim.ID)
	c.Set("isAdmin", claim.IsAdmin)
	c.Set("token", token)

	c.Next()
}

// Extract bearer token from Authorization header
func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad value authorization")
	}

	// Bearer schema only has template :
	// Bearer xxxRandomTOkenxxx. The authorization header has 2 element after splitted
	jwtToken := strings.Split(header, " ")

	if len(jwtToken) != 2 {
		return "", errors.New("invalid authorization format")
	}

	if jwtToken[0] != "Bearer" {
		return "", errors.New("invalid authorization scheme")
	}

	return jwtToken[1], nil
}

package jwt

import (
	"errors"
	"fmt"
	oauth "golang-bootcamp-1/internal/oauth/dto"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key []byte

var timeExpiration time.Time

// Generate JWT token for auth
// claim is custom claim to be inserted in JWT body
func GenerateToken(claim oauth.ClaimResponse) (string, *time.Time, error) {
	// Get key
	key = getKey()

	// Set time expiration
	timeExpiration = time.Now().Add(3 * 24 * time.Hour)

	// Set time expiration in JWT custom claim
	claim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(timeExpiration),
	}

	// Required variable for generating token
	var (
		t      *jwt.Token
		signed string
	)

	// Generate token
	t = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claim,
	)

	// Generate string / stringify token
	signed, err := t.SignedString(key)
	if err != nil {
		return "", nil, err
	}

	return signed, &timeExpiration, nil
}

func ValidateToken(token string) (*oauth.ClaimResponse, error) {
	// Get key
	key = getKey()

	parsed, err := jwt.ParseWithClaims(
		token,
		&oauth.ClaimResponse{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected token signing method : %v", t.Header["alg"])
			}

			return getKey(), nil
		},
	)
	if err != nil {
		return nil, err
	} else if claims, ok := parsed.Claims.(*oauth.ClaimResponse); ok {
		return claims, nil
	} else {
		return nil, errors.New("invalid claim type")
	}
}

func getKey() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

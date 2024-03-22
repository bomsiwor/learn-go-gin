package oauth

import "github.com/golang-jwt/jwt/v5"

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Type         string `json:"type"`
	ExpiredAt    string `json:"expiredAt"`
	Scope        string `json:"scope"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClaimResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin,omitempty"`
	jwt.RegisteredClaims
}

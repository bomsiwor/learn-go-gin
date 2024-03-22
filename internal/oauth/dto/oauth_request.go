package oauth

type LoginRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

type RefreshTokenRequest struct {
	Token string `json:"refreshToken"`
}

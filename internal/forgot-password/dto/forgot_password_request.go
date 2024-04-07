package forgot_password

import "time"

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ForgotPasswordUpdateRequest struct {
	Code     string `json:"code"`
	Password string `json:"password"`
}

type ForgotPasswordCreateRequest struct {
	UserID    int        `json:"userId"`
	Code      string     `json:"code"`
	Valid     bool       `json:"valid"`
	ExpiredAt *time.Time `json:"expiredAt"`
}

type ForgotPasswordEmailBody struct {
	Subject          string
	Email            string
	VerificationCode string
}

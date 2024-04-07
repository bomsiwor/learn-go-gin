package user

import "time"

type UserRequestBody struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedBy *int   `json:"createdBy"`
}

type UserUpdateRequest struct {
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Password        *string    `json:"password"`
	EmailVerifiedAt *time.Time `json:"emailVerifiedAt"`
	UpdatedBy       *int       `json:"updatedBy"`
}

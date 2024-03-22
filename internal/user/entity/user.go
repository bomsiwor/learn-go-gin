package user

import "time"

type User struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Password        string     `json:"-"`
	CodeVerified    string     `json:"-"`
	EmailVerifiedAt *time.Time `json:"verifiedAt"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
	DeletedAt       *time.Time `json:"deletedAt"`
}

package oauth

import (
	"time"

	"gorm.io/gorm"
)

type OauthClient struct {
	ID           int            `json:"id"`
	ClientID     string         `json:"clientId"`
	ClientSecret string         `json:"clientSecret"`
	Name         string         `json:"name"`
	Redirect     string         `json:"redirect"`
	Scope        string         `json:"scope"`
	CreatedAt    *time.Time     `json:"createdAt"`
	UpdatedAt    *time.Time     `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
}

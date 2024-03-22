package oauth

import (
	"time"

	"gorm.io/gorm"
)

type OauthAccessToken struct {
	ID            int            `json:"id"`
	UserID        int            `json:"userId"`
	OauthClientID *int           `json:"oauthClientId" gorm:"column:oauth_client_id"`
	OauthClient   *OauthClient   `gorm:"foreignKey:OauthClientID;references:ID"`
	Token         string         `json:"token"`
	Scope         string         `json:"scope"`
	ExpiredAt     *time.Time     `json:"expiredAt"`
	CreatedAt     *time.Time     `json:"createdBy"`
	UpdatedAt     *time.Time     `json:"updatedBy"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt"`
}

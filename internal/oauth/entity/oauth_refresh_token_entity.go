package oauth

import (
	"time"

	"gorm.io/gorm"
)

type OauthRefreshToken struct {
	ID                 int               `json:"id"`
	OauthAccessTokenID *int              `json:"oauthAccessTokenId" gorm:"column:oauth_access_token_id"`
	OauthAccessToken   *OauthAccessToken `gorm:"foreignKey:OauthAccessTokenID;references:ID"`
	UserId             int               `json:"userId"`
	Token              string            `json:"token"`
	ExpiredAt          *time.Time        `json:"expiredAt"`
	CreatedAt          *time.Time        `json:"createdBy"`
	UpdatedAt          *time.Time        `json:"updatedBy"`
	DeletedAt          gorm.DeletedAt    `json:"deletedAt"`
}

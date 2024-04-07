package forgot_password

import (
	user "golang-bootcamp-1/internal/user/entity"
	"time"

	"gorm.io/gorm"
)

type ForgotPassword struct {
	ID        int             `json:"id"`
	UserID    *int            `json:"userId"`
	User      *user.User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Valid     bool            `json:"valid"`
	Code      string          `json:"code"`
	ExpiredAt *time.Time      `json:"expiredAt"`
	CreatedAt *time.Time      `json:"createdAt"`
	UpdatedAt *time.Time      `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt"`
}

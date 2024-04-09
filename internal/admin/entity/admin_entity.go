package admin

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Password    string     `json:"-"`
	CreatedByID *int       `json:"-" gorm:"column:created_by"`
	CreatedBy   *Admin     `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int       `json:"-" gorm:"column:updated_by"`
	UpdatedBy   *Admin     `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt
}

package product_category

import (
	admin "golang-bootcamp-1/internal/admin/entity"
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID          int64        `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"size:255;not null" json:"name"`
	Image       string       `gorm:"size:255" json:"image" `
	CreatedById *int64       `gorm:"not null;column:created_by" json:"createdById"`
	CreatedBy   *admin.Admin `json:"createdBy" gorm:"foreignKey:CreatedById;references:ID"`
	UpdatedById *int64       `json:"updatedById"`
	UpdatedBy   *admin.Admin `json:"updatedBy" gorm:"foreignKey:UpdatedById;references:ID"`
	CreatedAt   *time.Time   `json:"createdAt"`
	UpdatedAt   *time.Time   `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt
}

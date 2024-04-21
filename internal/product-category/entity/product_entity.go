package product_category

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                int              `gorm:"primaryKey" json:"id"`
	ProductCategoryID *int             `json:"productCategoryId" gorm:"column:product_category_id"`
	Category          *ProductCategory `json:"category" gorm:"foreignKey:ProductCategoryID;references:ID"`
	Title             string           `gorm:"size:255" json:"title"`
	Image             *string          `gorm:"size:255" json:"image"`
	Video             *string          `gorm:"size:255" json:"video"`
	Description       *string          `gorm:"type:text" json:"description"`
	IsHighlighted     bool             `json:"isHighlighted"`
	Price             float64          `gorm:"type:decimal(10,2)" json:"price"`
	CreatedById       *int             `json:"createdById"`
	UpdatedById       *int             `json:"updatedById"`
	CreatedAt         *time.Time       `json:"createdAt"`
	UpdatedAt         *time.Time       `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt   `json:"-"`
}

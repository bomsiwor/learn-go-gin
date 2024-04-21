package product_category

import "mime/multipart"

type ProducteCategoryRequest struct {
	Name      string                `form:"name"`
	Image     *multipart.FileHeader `form:"image"`
	CreatedBy *int64
	UpdatedBy *int64
}

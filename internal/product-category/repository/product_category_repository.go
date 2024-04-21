package product_category

import (
	"errors"
	entity "golang-bootcamp-1/internal/product-category/entity"
	"golang-bootcamp-1/pkg/response"
	"golang-bootcamp-1/pkg/utils"

	"gorm.io/gorm"
)

type IProductCategory interface {
	FindAll(page, limit int) []entity.ProductCategory
	FindById(id int) (*entity.ProductCategory, *response.ErrorResp)
	Create(entity entity.ProductCategory) (*entity.ProductCategory, *response.ErrorResp)
	Update(entity entity.ProductCategory, id int) (*entity.ProductCategory, *response.ErrorResp)
	Delete(entity entity.ProductCategory) *response.ErrorResp
}

type productCategoryRepository struct {
	db *gorm.DB
}

// Create implements IProductCategory.
func (repo *productCategoryRepository) Create(entity entity.ProductCategory) (*entity.ProductCategory, *response.ErrorResp) {
	if err := repo.db.Create(&entity).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &entity, nil
}

// Delete implements IProductCategory.
func (repo *productCategoryRepository) Delete(entity entity.ProductCategory) *response.ErrorResp {
	if err := repo.db.Delete(&entity).Error; err != nil {
		return &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return nil
}

// FindAll implements IProductCategory.
func (repo *productCategoryRepository) FindAll(page int, limit int) []entity.ProductCategory {
	var productCategories []entity.ProductCategory

	repo.db.Scopes(utils.Paginate(page, limit)).Find(&productCategories)

	return productCategories
}

// FindById implements IProductCategory.
func (repo *productCategoryRepository) FindById(id int) (*entity.ProductCategory, *response.ErrorResp) {
	var productCategory entity.ProductCategory

	if err := repo.db.Where("id = ?", id).First(&productCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResp{
				Code:    404,
				Err:     err,
				Message: "Not found",
			}
		}

		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: "Error retrieving data",
		}
	}

	return &productCategory, nil
}

// Update implements IProductCategory.
func (repo *productCategoryRepository) Update(entity entity.ProductCategory, id int) (*entity.ProductCategory, *response.ErrorResp) {
	if err := repo.db.Where("id=?", id).Save(&entity).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &entity, nil
}

func NewProductCategoryRepository(db *gorm.DB) IProductCategory {
	return &productCategoryRepository{
		db: db,
	}
}

package product_category

import (
	dto "golang-bootcamp-1/internal/product-category/dto"
	entity "golang-bootcamp-1/internal/product-category/entity"
	repo "golang-bootcamp-1/internal/product-category/repository"
	"golang-bootcamp-1/pkg/media"
	"golang-bootcamp-1/pkg/response"
)

type IProductCategoryUsecase interface {
	FindAll(page, limit int) []entity.ProductCategory
	FindById(id int) (*entity.ProductCategory, *response.ErrorResp)
	Create(dto dto.ProducteCategoryRequest) (*entity.ProductCategory, *response.ErrorResp)
	Update(dto dto.ProducteCategoryRequest, id int) (*entity.ProductCategory, *response.ErrorResp)
	Delete(id int) *response.ErrorResp
}

type productCategoryUsecase struct {
	repo  repo.IProductCategory
	media media.IMedia
}

func NewProductCategoryUsecase(repo repo.IProductCategory, media media.IMedia) IProductCategoryUsecase {
	return &productCategoryUsecase{
		repo:  repo,
		media: media,
	}
}

// Create implements IProductCategoryUsecase.
func (uc *productCategoryUsecase) Create(request dto.ProducteCategoryRequest) (*entity.ProductCategory, *response.ErrorResp) {
	productCategory := entity.ProductCategory{
		Name:        request.Name,
		CreatedById: request.CreatedBy,
	}

	// Upload image
	if request.Image != nil {
		image, err := uc.media.Upload(request.Image, "category")
		if err != nil {
			return nil, err
		}
		// If upload succeeded, store url
		productCategory.Image = image
	}

	// Store to database
	// Do not forget to customize error message
	return uc.repo.Create(productCategory)
}

// Delete implements IProductCategoryUsecase.
func (uc *productCategoryUsecase) Delete(id int) *response.ErrorResp {
	// Find data by ID
	productCategory, err := uc.repo.FindById(id)
	if err != nil {
		return err
	}

	// Delete data
	return uc.repo.Delete(*productCategory)
}

// FindAll implements IProductCategoryUsecase.
func (uc *productCategoryUsecase) FindAll(page int, limit int) []entity.ProductCategory {
	return uc.repo.FindAll(page, limit)
}

// FindById implements IProductCategoryUsecase.
func (uc *productCategoryUsecase) FindById(id int) (*entity.ProductCategory, *response.ErrorResp) {
	productCategory, err := uc.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

// Update implements IProductCategoryUsecase.
func (uc *productCategoryUsecase) Update(dto dto.ProducteCategoryRequest, id int) (*entity.ProductCategory, *response.ErrorResp) {
	panic("unimplemented")
}

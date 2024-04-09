package admin

import (
	"errors"
	entity "golang-bootcamp-1/internal/admin/entity"
	"golang-bootcamp-1/pkg/response"
	"golang-bootcamp-1/pkg/utils"

	"gorm.io/gorm"
)

type IAdminRepository interface {
	FindAll(page, limit int) []entity.Admin
	FindByID(id int) (*entity.Admin, *response.ErrorResp)
	FindByEmail(email string) (*entity.Admin, *response.ErrorResp)
	Create(entity entity.Admin) (*entity.Admin, *response.ErrorResp)
	Update(entity entity.Admin) (*entity.Admin, *response.ErrorResp)
	Delete(entity entity.Admin) *response.ErrorResp
	TotalCountAdmin() (int64, *response.ErrorResp)
}

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) IAdminRepository {
	return &AdminRepository{
		db: db,
	}
}

// Create implements IAdminRepository.
func (repo *AdminRepository) Create(entity entity.Admin) (*entity.Admin, *response.ErrorResp) {
	if err := repo.db.Create(&entity).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &entity, nil
}

// Delete implements IAdminRepository.
func (repo *AdminRepository) Delete(entity entity.Admin) *response.ErrorResp {
	if err := repo.db.Delete(&entity).Error; err != nil {
		return &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return nil
}

// FindAll implements IAdminRepository.
func (repo *AdminRepository) FindAll(page int, limit int) []entity.Admin {
	var admins []entity.Admin

	repo.db.Scopes(utils.Paginate(page, limit)).Find(&admins)

	return admins
}

// FindByEmail implements IAdminRepository.
func (repo *AdminRepository) FindByEmail(email string) (*entity.Admin, *response.ErrorResp) {
	var admin entity.Admin

	if err := repo.db.Where("email = ?", email).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResp{
				Code:    404,
				Err:     err,
				Message: "Data tidak ditemukan",
			}
		}

		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &admin, nil
}

// FindByID implements IAdminRepository.
func (repo *AdminRepository) FindByID(id int) (*entity.Admin, *response.ErrorResp) {
	var admin entity.Admin

	if err := repo.db.Where("id = ?", id).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResp{
				Code:    404,
				Err:     err,
				Message: "Data tidak ditemukan",
			}
		}

		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &admin, nil
}

// TotalCountAdmin implements IAdminRepository.
func (repo *AdminRepository) TotalCountAdmin() (int64, *response.ErrorResp) {
	var count int64

	if err := repo.db.Model(&entity.Admin{}).Count(&count).Error; err != nil {
		return 0, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return count, nil
}

// Update implements IAdminRepository.
func (repo *AdminRepository) Update(entity entity.Admin) (*entity.Admin, *response.ErrorResp) {
	if err := repo.db.Save(&entity).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &entity, nil
}

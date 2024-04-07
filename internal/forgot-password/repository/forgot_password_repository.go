package forgot_password

import (
	"errors"
	entity "golang-bootcamp-1/internal/forgot-password/entity"
	"golang-bootcamp-1/pkg/response"

	"gorm.io/gorm"
)

type IForgotPasswordRepo interface {
	Create(entity entity.ForgotPassword) (*entity.ForgotPassword, *response.ErrorResp)
	FindByCode(code string) (*entity.ForgotPassword, *response.ErrorResp)
	Update(entity entity.ForgotPassword) (*entity.ForgotPassword, *response.ErrorResp)
}

type forgotPasswordRepo struct {
	db *gorm.DB
}

func NewForgotPasswordRepository(db *gorm.DB) IForgotPasswordRepo {
	return &forgotPasswordRepo{
		db: db,
	}
}

// Create implements IForgotPasswordRepo.
func (repo *forgotPasswordRepo) Create(entity entity.ForgotPassword) (*entity.ForgotPassword, *response.ErrorResp) {
	if err := repo.db.Create(&entity).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &entity, nil
}

// Find data by unique code
func (repo *forgotPasswordRepo) FindByCode(code string) (*entity.ForgotPassword, *response.ErrorResp) {
	// Data to store found data
	var data entity.ForgotPassword

	// Query data and check the error
	// Error can be not found
	if err := repo.db.Where("code = ?", code).First(&data).Error; err != nil {
		// IF error is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResp{
				Code:    404,
				Err:     err,
				Message: "Code not valid",
			}
		}

		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &data, nil
}

// Update the data and return new data
func (repo *forgotPasswordRepo) Update(entity entity.ForgotPassword) (*entity.ForgotPassword, *response.ErrorResp) {
	if err := repo.db.Save(&entity).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &entity, nil
}

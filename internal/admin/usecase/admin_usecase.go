package admin

import (
	"errors"
	dto "golang-bootcamp-1/internal/admin/dto"
	entity "golang-bootcamp-1/internal/admin/entity"
	repo "golang-bootcamp-1/internal/admin/repository"
	"golang-bootcamp-1/pkg/hasher"
	"golang-bootcamp-1/pkg/response"
	"net/http"
)

type IAdminUsecase interface {
	FindAll(page, limit int) []entity.Admin
	FindByID(id int) (*entity.Admin, *response.ErrorResp)
	FindByEmail(email string) (*entity.Admin, *response.ErrorResp)
	Create(dto dto.AdminRequestBody) (*entity.Admin, *response.ErrorResp)
	Update(id int, dto dto.AdminRequestBody) (*entity.Admin, *response.ErrorResp)
	Delete(id int) *response.ErrorResp
	TotalCountAdmin() (int64, *response.ErrorResp)
}

type AdminUsecase struct {
	repo repo.IAdminRepository
}

func NewAdminUsecase(repo repo.IAdminRepository) IAdminUsecase {
	return &AdminUsecase{
		repo: repo,
	}
}

// Create implements IAdminUsecase.
func (uc *AdminUsecase) Create(dto dto.AdminRequestBody) (*entity.Admin, *response.ErrorResp) {
	// Check if data already exists by email
	prev, _ := uc.repo.FindByEmail(dto.Email)

	// If data exists, cancel process
	if prev != nil {
		return nil, &response.ErrorResp{
			Code:    http.StatusConflict,
			Message: "Data already exists",
			Err:     errors.New(http.StatusText(http.StatusConflict)),
		}
	}

	// Create admin data
	// Generate password
	hashedPassword, err := hasher.GeneratePassword(*dto.Password)
	if err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: "Failed to process request",
		}
	}

	// Wrap in model
	data := entity.Admin{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashedPassword,
	}

	// Save in db via repo
	result, errCreate := uc.repo.Create(data)
	if err != nil {
		return nil, &response.ErrorResp{
			Code:    errCreate.Code,
			Err:     errCreate.Err,
			Message: errCreate.Message,
		}
	}

	return result, nil
}

// Delete implements IAdminUsecase.
func (uc *AdminUsecase) Delete(id int) *response.ErrorResp {
	// Search by ID
	admin, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Delete data
	return uc.repo.Delete(*admin)
}

// FindAll implements IAdminUsecase.
func (uc *AdminUsecase) FindAll(page int, limit int) []entity.Admin {
	return uc.repo.FindAll(page, limit)
}

// FindByEmail implements IAdminUsecase.
func (uc *AdminUsecase) FindByEmail(email string) (*entity.Admin, *response.ErrorResp) {
	return uc.repo.FindByEmail(email)
}

// FindByID implements IAdminUsecase.
func (uc *AdminUsecase) FindByID(id int) (*entity.Admin, *response.ErrorResp) {
	return uc.repo.FindByID(id)
}

// TotalCountAdmin implements IAdminUsecase.
func (uc *AdminUsecase) TotalCountAdmin() (int64, *response.ErrorResp) {
	return uc.repo.TotalCountAdmin()
}

// Update implements IAdminUsecase.
func (uc *AdminUsecase) Update(id int, dto dto.AdminRequestBody) (*entity.Admin, *response.ErrorResp) {
	// Find data by id
	admin, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update admin query result
	admin.Name = dto.Name
	admin.Email = dto.Email

	// Generate new password if updated
	if dto.Password != nil {
		hashedPassword, errHash := hasher.GeneratePassword(*dto.Password)
		if errHash != nil {
			return nil, &response.ErrorResp{
				Code:    500,
				Err:     errHash,
				Message: "Failed to update data",
			}
		}

		admin.Password = hashedPassword
	}

	return uc.repo.Update(*admin)
}

package usecase

import (
	"errors"
	dto "golang-bootcamp-1/internal/user/dto"
	entity "golang-bootcamp-1/internal/user/entity"
	repository "golang-bootcamp-1/internal/user/repository"
	"golang-bootcamp-1/pkg/hasher"
	"golang-bootcamp-1/pkg/response"
	"golang-bootcamp-1/pkg/utils"

	"gorm.io/gorm"
)

type IUserUseCase interface {
	FindAll(offset, limit int) []entity.User
	FindById(id int) (*entity.User, *response.ErrorResp)
	FindByEmail(email string) (*entity.User, *response.ErrorResp)
	Create(dto dto.UserRequestBody) (*entity.User, *response.ErrorResp)
	Delete(id int) *response.ErrorResp
}

type userUseCase struct {
	repository repository.IUserRepo
}

// Create implements IUserUseCase.
func (usecase *userUseCase) Create(dto dto.UserRequestBody) (*entity.User, *response.ErrorResp) {
	// Check by email
	checkUser, err := usecase.repository.FindByEmail(dto.Email)

	// Check if any error and not error not found
	if err != nil && !errors.Is(err.Err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// If user exists - dont proceed further
	if checkUser != nil {
		return nil, &response.ErrorResp{
			Code:    409,
			Message: "User exists!",
			Err:     errors.New("user exists"),
		}
	}

	// Create User
	// Generate password
	password, errHash := hasher.GeneratePassword(dto.Password)
	if errHash != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Message: "Error processing",
			Err:     errors.New("error processing"),
		}
	}
	// Put into user entity
	var newUser entity.User = entity.User{
		Name:         dto.Name,
		Email:        dto.Email,
		Password:     password,
		CodeVerified: utils.RandomString(32),
	}

	userData, err := usecase.repository.Create(newUser)
	if err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Message: "Error processing",
			Err:     errors.New("error processing"),
		}
	}

	return userData, nil
}

// Delete implements IUserUseCase.
func (usecase *userUseCase) Delete(id int) *response.ErrorResp {
	panic("unimplemented")
}

// FindAll implements IUserUseCase.
func (usecase *userUseCase) FindAll(offset int, limit int) []entity.User {
	panic("unimplemented")
}

// FindByEmail implements IUserUseCase.
func (usecase *userUseCase) FindByEmail(email string) (*entity.User, *response.ErrorResp) {
	return usecase.repository.FindByEmail(email)
}

// FindById implements IUserUseCase.
func (usecase *userUseCase) FindById(id int) (*entity.User, *response.ErrorResp) {
	panic("unimplemented")
}

func NewUserUseCase(repo repository.IUserRepo) IUserUseCase {
	return &userUseCase{
		repository: repo,
	}
}

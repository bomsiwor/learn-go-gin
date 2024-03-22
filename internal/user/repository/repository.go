package user

import (
	"errors"
	entity "golang-bootcamp-1/internal/user/entity"
	"golang-bootcamp-1/pkg/response"

	"gorm.io/gorm"
)

type IUserRepo interface {
	FindAll(offset, limit int) []entity.User
	FindById(id int) (*entity.User, *response.ErrorResp)
	FindByEmail(email string) (*entity.User, *response.ErrorResp)
	Create(entity entity.User) (*entity.User, *response.ErrorResp)
	Delete(id int) *response.ErrorResp
}

type userRepository struct {
	db *gorm.DB
}

// Create implements IUserRepo.
func (repo *userRepository) Create(entity entity.User) (*entity.User, *response.ErrorResp) {
	if err := repo.db.Create(&entity).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return &entity, nil
}

// Delete implements IUserRepo.
func (repo *userRepository) Delete(id int) *response.ErrorResp {
	panic("unimplemented")
}

// FindAll implements IUserRepo.
func (repo *userRepository) FindAll(offset int, limit int) []entity.User {
	var users []entity.User

	repo.db.Limit(limit).Offset(offset).Find(&users)

	return users
}

// FindByEmail implements IUserRepo.
func (repo *userRepository) FindByEmail(email string) (*entity.User, *response.ErrorResp) {
	var user entity.User
	userData := repo.db.Where("email = ?", email).First(&user)

	// If error query
	if err := userData.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResp{
				Code:    404,
				Message: err.Error(),
				Err:     err,
			}
		}

		return nil, &response.ErrorResp{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return &user, nil
}

// FindById implements IUserRepo.
func (repo *userRepository) FindById(id int) (*entity.User, *response.ErrorResp) {
	var user entity.User

	userData := repo.db.Where("id = ?", id).First(&user)
	if err := userData.Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return &user, nil
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepository{
		db: db,
	}
}

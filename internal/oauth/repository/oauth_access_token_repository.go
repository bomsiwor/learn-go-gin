package oauth

import (
	entity "golang-bootcamp-1/internal/oauth/entity"
	"golang-bootcamp-1/pkg/response"

	"gorm.io/gorm"
)

type IOauthAccessTokenRepo interface {
	FindByAccessToken(token string) (*entity.OauthAccessToken, *response.ErrorResp)
	Delete(entity entity.OauthAccessToken) *response.ErrorResp
	Create(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.ErrorResp)
}

type oauthAccessTokenRepo struct {
	db *gorm.DB
}

// Create implements IOauthAccessTokenRepo.
func (repo *oauthAccessTokenRepo) Create(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.ErrorResp) {
	// Create data
	if err := repo.db.Create(&entity).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &entity, nil
}

// Delete implements IOauthAccessTokenRepo.
func (repo *oauthAccessTokenRepo) Delete(entity entity.OauthAccessToken) *response.ErrorResp {
	if err := repo.db.Delete(&entity).Error; err != nil {
		return &response.ErrorResp{
			Code:    500,
			Message: "Failed to invalidate token",
			Err:     err,
		}
	}

	return nil
}

// FindByAccessToken implements IOauthAccessTokenRepo.
func (repo *oauthAccessTokenRepo) FindByAccessToken(token string) (*entity.OauthAccessToken, *response.ErrorResp) {
	panic("unimplemented")
}

func NewOauthAcces(db *gorm.DB) IOauthAccessTokenRepo {
	return &oauthAccessTokenRepo{
		db: db,
	}
}

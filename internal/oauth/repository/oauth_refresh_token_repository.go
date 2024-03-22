package oauth

import (
	entity "golang-bootcamp-1/internal/oauth/entity"
	"golang-bootcamp-1/pkg/response"

	"gorm.io/gorm"
)

type IOauthRefreshTokenRepo interface {
	Create(entity entity.OauthRefreshToken) (*entity.OauthRefreshToken, *response.ErrorResp)
	FindByToken(token string) (*entity.OauthRefreshToken, *response.ErrorResp)
	FindByOauthAccessTokenID(accessTokenId int) (*entity.OauthRefreshToken, *response.ErrorResp)
	Delete(entity entity.OauthRefreshToken) *response.ErrorResp
}

type oauthRefreshTokenRepo struct {
	db *gorm.DB
}

// Create implements IOauthRefreshTokenRepo.
func (repo *oauthRefreshTokenRepo) Create(entity entity.OauthRefreshToken) (*entity.OauthRefreshToken, *response.ErrorResp) {
	if err := repo.db.Create(&entity).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &entity, nil
}

// Delete implements IOauthRefreshTokenRepo.
func (repo *oauthRefreshTokenRepo) Delete(entity entity.OauthRefreshToken) *response.ErrorResp {
	panic("unimplemented")
}

// FindByOauthAccessTokenID implements IOauthRefreshTokenRepo.
func (repo *oauthRefreshTokenRepo) FindByOauthAccessTokenID(accessTokenId int) (*entity.OauthRefreshToken, *response.ErrorResp) {
	panic("unimplemented")
}

// FindByToken implements IOauthRefreshTokenRepo.
func (repo *oauthRefreshTokenRepo) FindByToken(token string) (*entity.OauthRefreshToken, *response.ErrorResp) {
	panic("unimplemented")
}

func NewOauthRefreshTokenRepo(db *gorm.DB) IOauthRefreshTokenRepo {
	return &oauthRefreshTokenRepo{
		db: db,
	}
}

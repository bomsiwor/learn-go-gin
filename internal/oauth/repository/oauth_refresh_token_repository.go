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
	err := repo.db.Delete(&entity).Error
	if err != nil {
		return &response.ErrorResp{
			Code:    500,
			Message: "Failed to invalidate token",
			Err:     err,
		}
	}

	return nil
}

// FindByOauthAccessTokenID implements IOauthRefreshTokenRepo.
func (repo *oauthRefreshTokenRepo) FindByOauthAccessTokenID(accessTokenId int) (*entity.OauthRefreshToken, *response.ErrorResp) {
	var refreshToken entity.OauthRefreshToken

	if err := repo.db.Where("oauth_access_token_id = ?", accessTokenId).First(&refreshToken).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    401,
			Err:     err,
			Message: "Token invalid",
		}
	}

	return &refreshToken, nil
}

// FindByToken implements IOauthRefreshTokenRepo.
func (repo *oauthRefreshTokenRepo) FindByToken(token string) (*entity.OauthRefreshToken, *response.ErrorResp) {
	var refreshToken entity.OauthRefreshToken

	if err := repo.db.Where("token =?", token).Preload("OauthAccessToken.OauthClient").First(&refreshToken).Error; err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: "Refresh token invalid",
		}
	}

	return &refreshToken, nil
}

func NewOauthRefreshTokenRepo(db *gorm.DB) IOauthRefreshTokenRepo {
	return &oauthRefreshTokenRepo{
		db: db,
	}
}

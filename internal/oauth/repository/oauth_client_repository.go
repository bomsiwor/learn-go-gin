package oauth

import (
	entity "golang-bootcamp-1/internal/oauth/entity"
	"golang-bootcamp-1/pkg/response"

	"gorm.io/gorm"
)

type IOauthClientRepo interface {
	FindByClientIdAndClientSecret(clientId, clientSecret string) (*entity.OauthClient, *response.ErrorResp)
}

type oauthClientRepo struct {
	db *gorm.DB
}

// FindByClientIdAndClientSecret implements IOauthClientRepo.
// Search for client secret
func (repo *oauthClientRepo) FindByClientIdAndClientSecret(clientId string, clientSecret string) (*entity.OauthClient, *response.ErrorResp) {
	var oauthClient entity.OauthClient

	// Search by query
	// Return error if not found
	// If need a special treatment, handle on use case. So repository only interact with DB and returning data
	err := repo.db.Where("client_id = ?", clientId).Where("client_secret = ?", clientSecret).First(&oauthClient).Error
	if err != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err,
			Message: err.Error(),
		}
	}

	return &oauthClient, nil
}

func NewOauthClientRepo(db *gorm.DB) IOauthClientRepo {
	return &oauthClientRepo{
		db: db,
	}
}

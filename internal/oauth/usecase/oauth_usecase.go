package oauth

import (
	"errors"
	adminUsecase "golang-bootcamp-1/internal/admin/usecase"
	dto "golang-bootcamp-1/internal/oauth/dto"
	entity "golang-bootcamp-1/internal/oauth/entity"
	repository "golang-bootcamp-1/internal/oauth/repository"
	userDto "golang-bootcamp-1/internal/user/dto"
	userUsecase "golang-bootcamp-1/internal/user/usecase"
	"golang-bootcamp-1/pkg/hasher"
	"golang-bootcamp-1/pkg/jwt"
	"golang-bootcamp-1/pkg/response"
	"golang-bootcamp-1/pkg/utils"
	"time"
)

type IOauthUseCase interface {
	Login(request dto.LoginRequest) (*dto.LoginResponse, *response.ErrorResp)
	Logout(token string) *response.ErrorResp
	Refresh(request dto.RefreshTokenRequest) (*dto.LoginResponse, *response.ErrorResp)
	Me(id int, isAdmin bool) (*userDto.UserMeResponse, *response.ErrorResp)
}

type oauthUseCase struct {
	oauthClient       repository.IOauthClientRepo
	oauthAccessToken  repository.IOauthAccessTokenRepo
	oauthRefreshToken repository.IOauthRefreshTokenRepo
	userUsecase       userUsecase.IUserUseCase
	adminUsecase      adminUsecase.IAdminUsecase
}

// Logout implements IOauthUseCase.
func (uc *oauthUseCase) Logout(token string) *response.ErrorResp {
	// Search access token data by token
	// Delete after date retrieved
	accessToken, err := uc.oauthAccessToken.FindByAccessToken(token)
	if err != nil {
		return err
	}

	// Delete access token
	err = uc.oauthAccessToken.Delete(*accessToken)
	if err != nil {
		return err
	}

	// Delete refresh token
	refreshToken, err := uc.oauthRefreshToken.FindByOauthAccessTokenID(accessToken.ID)
	if err != nil {
		return err
	}
	err = uc.oauthRefreshToken.Delete(*refreshToken)
	if err != nil {
		return err
	}

	return nil
}

// Get current user login data
// This function is  on authentication context
func (uc *oauthUseCase) Me(id int, isAdmin bool) (*userDto.UserMeResponse, *response.ErrorResp) {
	var meData userDto.UserMeResponse

	if isAdmin {
		user, err := uc.adminUsecase.FindByID(id)
		if err != nil {
			return nil, &response.ErrorResp{
				Code:    409,
				Message: "User not found or invalid",
				Err:     errors.New("user not found"),
			}
		}
		meData.ID = user.ID
		meData.Email = user.Email
		meData.Name = user.Name
	} else {
		user, err := uc.userUsecase.FindById(id)
		if err != nil {
			return nil, &response.ErrorResp{
				Code:    409,
				Message: "User not found or invalid",
				Err:     errors.New("user not found"),
			}
		}
		meData.ID = user.ID
		meData.Email = user.Email
		meData.Name = user.Name
	}

	return &meData, nil
}

// Login implements IOauthUseCase.
func (uc *oauthUseCase) Login(request dto.LoginRequest) (*dto.LoginResponse, *response.ErrorResp) {
	// Check wheter client ID & secret is valid
	oauthClient, err := uc.oauthClient.FindByClientIdAndClientSecret(request.ClientID, request.ClientSecret)
	if err != nil {
		return nil, err
	}

	// Check user, wheter it is admin or user
	var user dto.UserResponse
	if oauthClient.Name == "web-admin" {
		// Check admin data
		adminData, err := uc.adminUsecase.FindByEmail(request.Email)
		if err != nil {
			return nil, err
		}

		// Do not mutate data from repo / query
		// Copy it on another DTO
		user.ID = adminData.ID
		user.Email = adminData.Email
		user.Name = adminData.Name
		user.Password = adminData.Password
	} else {
		// Check user data
		userData, err := uc.userUsecase.FindByEmail(request.Email)
		if err != nil {
			return nil, err
		}

		// Do not mutate data from repo / query
		// Copy it on another DTO
		user.ID = userData.ID
		user.Email = userData.Email
		user.Name = userData.Name
		user.Password = userData.Password
	}

	// Check password
	// Dont proceed if passwrod wrong
	errPassword := hasher.ValidatePassword(request.Password, user.Password)
	if errPassword != nil {
		return nil, &response.ErrorResp{
			Code:    409,
			Err:     errors.New("wrong credential"),
			Message: "Access Denied",
		}
	}

	// If password is correct, processs further
	customClaim := dto.ClaimResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
	}

	// If admin login, set isAdmin flag
	if oauthClient.Name == "web-admin" {
		customClaim.IsAdmin = true
	}

	token, expiredAt, errToken := jwt.GenerateToken(customClaim)
	if errToken != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     errToken,
			Message: "Failed to authenticate",
		}
	}

	// After token generated
	// Insert into oauth acccess token & refresh token
	// Insert oauth access token first
	oauthAccessTokenData := entity.OauthAccessToken{
		UserID:        user.ID,
		OauthClientID: &oauthClient.ID,
		Token:         token,
		ExpiredAt:     expiredAt,
		Scope:         "*",
	}

	oauthAccessToken, errAccessToken := uc.oauthAccessToken.Create(oauthAccessTokenData)
	if errAccessToken != nil {
		return nil, errAccessToken
	}

	// Insert into refresh token
	refreshExpiration := time.Now().Add(7 * 24 * time.Hour)
	refreshTokenData := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		Token:              utils.RandomString(128),
		ExpiredAt:          &refreshExpiration,
		UserId:             user.ID,
	}

	refreshToken, err := uc.oauthRefreshToken.Create(refreshTokenData)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  oauthAccessToken.Token,
		RefreshToken: refreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expiredAt.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

// Refresh implements IOauthUseCase.
func (uc *oauthUseCase) Refresh(request dto.RefreshTokenRequest) (*dto.LoginResponse, *response.ErrorResp) {
	// Check for refresh token
	refreshToken, err := uc.oauthRefreshToken.FindByToken(request.Token)
	if err != nil {
		return nil, err
	}

	// Check for time expiration
	if refreshToken.ExpiredAt.Before(time.Now().UTC()) {
		return nil, &response.ErrorResp{
			Code:    401,
			Message: "Token expired",
			Err:     errors.New("token expired"),
		}
	}

	// If token valid
	// Invalidate old refresh token by deleting it
	// Invalidate old access token by deleting
	// Generate new access token and refresh token
	clientId := refreshToken.OauthAccessToken.OauthClientID
	userId := refreshToken.UserId

	err = uc.oauthAccessToken.Delete(*refreshToken.OauthAccessToken)
	if err != nil {
		return nil, err
	}

	err = uc.oauthRefreshToken.Delete(*refreshToken)
	if err != nil {
		return nil, err
	}

	// Regenerate token
	var user dto.UserResponse
	if refreshToken.OauthAccessToken.OauthClient.Name == "web-admin" {
		// Check admin data
		adminData, err := uc.adminUsecase.FindByID(userId)
		if err != nil {
			return nil, err
		}

		// Do not mutate data from repo / query
		// Copy it on another DTO
		user.ID = adminData.ID
		user.Email = adminData.Email
		user.Name = adminData.Name
		user.Password = adminData.Password
	} else {
		// Check user data
		userData, err := uc.userUsecase.FindById(userId)
		if err != nil {
			return nil, err
		}

		// Do not mutate data from repo / query
		// Copy it on another DTO
		user.ID = userData.ID
		user.Email = userData.Email
		user.Name = userData.Name
		user.Password = userData.Password
	}

	customClaim := dto.ClaimResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
	}

	token, expiredAt, errToken := jwt.GenerateToken(customClaim)
	if errToken != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     errToken,
			Message: "Failed to authenticate",
		}
	}

	// After token generated
	// Insert into oauth acccess token & refresh token
	// Insert oauth access token first
	oauthAccessTokenData := entity.OauthAccessToken{
		UserID:        user.ID,
		OauthClientID: clientId,
		Token:         token,
		ExpiredAt:     expiredAt,
		Scope:         "*",
	}

	oauthAccessToken, errAccessToken := uc.oauthAccessToken.Create(oauthAccessTokenData)
	if errAccessToken != nil {
		return nil, &response.ErrorResp{
			Code:    500,
			Err:     err.Err,
			Message: err.Err.Error(),
		}
	}

	// Insert into refresh token
	refreshExpiration := time.Now().Add(7 * 24 * time.Hour)
	refreshTokenData := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		Token:              utils.RandomString(128),
		ExpiredAt:          &refreshExpiration,
		UserId:             user.ID,
	}

	refreshToken, err = uc.oauthRefreshToken.Create(refreshTokenData)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  oauthAccessToken.Token,
		RefreshToken: refreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expiredAt.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

func NewOauthUseCase(
	client repository.IOauthClientRepo,
	accessToken repository.IOauthAccessTokenRepo,
	refreshToken repository.IOauthRefreshTokenRepo,
	userUseCase userUsecase.IUserUseCase,
	adminUsecase adminUsecase.IAdminUsecase,
) IOauthUseCase {
	return &oauthUseCase{
		oauthClient:       client,
		oauthAccessToken:  accessToken,
		oauthRefreshToken: refreshToken,
		userUsecase:       userUseCase,
		adminUsecase:      adminUsecase,
	}
}

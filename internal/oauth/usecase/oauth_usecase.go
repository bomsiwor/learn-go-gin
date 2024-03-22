package oauth

import (
	"errors"
	"fmt"
	dto "golang-bootcamp-1/internal/oauth/dto"
	entity "golang-bootcamp-1/internal/oauth/entity"
	repository "golang-bootcamp-1/internal/oauth/repository"
	userUsecase "golang-bootcamp-1/internal/user/usecase"
	"golang-bootcamp-1/pkg/hasher"
	"golang-bootcamp-1/pkg/response"
	"golang-bootcamp-1/pkg/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IOauthUseCase interface {
	Login(request dto.LoginRequest) (*dto.LoginResponse, *response.ErrorResp)
}

type oauthUseCase struct {
	oauthClient       repository.IOauthClientRepo
	oauthAccessToken  repository.IOauthAccessTokenRepo
	oauthRefreshToken repository.IOauthRefreshTokenRepo
	userUsecase       userUsecase.IUserUseCase
}

// Login implements IOauthUseCase.
func (uc *oauthUseCase) Login(request dto.LoginRequest) (*dto.LoginResponse, *response.ErrorResp) {
	// Check wheter client ID & secret is valid
	fmt.Println(request)
	oauthClient, err := uc.oauthClient.FindByClientIdAndClientSecret(request.ClientID, request.ClientSecret)
	if err != nil {
		return nil, err
	}

	// Check user data
	var user dto.UserResponse

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
	// Set JWT Expiration
	timeExpiration := time.Now().Add(3 * 24 * time.Hour)

	// Set variable required for token generation
	var (
		key         []byte
		t           *jwt.Token
		signed      string
		customClaim dto.ClaimResponse
	)

	customClaim = dto.ClaimResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(timeExpiration),
		},
	}

	key = []byte(os.Getenv("JWT_SECRET_KEY"))
	t = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		customClaim,
	)

	signed, _ = t.SignedString(key)

	// After token generated
	// Insert into oauth acccess token & refresh token
	// Insert oauth access token first
	oauthAccessTokenData := entity.OauthAccessToken{
		UserID:        user.ID,
		OauthClientID: &oauthClient.ID,
		Token:         signed,
		ExpiredAt:     &timeExpiration,
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

	refreshToken, err := uc.oauthRefreshToken.Create(refreshTokenData)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  oauthAccessToken.Token,
		RefreshToken: refreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    timeExpiration.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

func NewOauthUseCase(
	client repository.IOauthClientRepo,
	accessToken repository.IOauthAccessTokenRepo,
	refreshToken repository.IOauthRefreshTokenRepo,
	userUseCase userUsecase.IUserUseCase,
) IOauthUseCase {
	return &oauthUseCase{
		oauthClient:       client,
		oauthAccessToken:  accessToken,
		oauthRefreshToken: refreshToken,
		userUsecase:       userUseCase,
	}
}

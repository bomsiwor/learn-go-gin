package forgot_password

import (
	"errors"
	dto "golang-bootcamp-1/internal/forgot-password/dto"
	entity "golang-bootcamp-1/internal/forgot-password/entity"
	repo "golang-bootcamp-1/internal/forgot-password/repository"
	userDto "golang-bootcamp-1/internal/user/dto"
	userUc "golang-bootcamp-1/internal/user/usecase"
	mail "golang-bootcamp-1/pkg/mail"

	"golang-bootcamp-1/pkg/response"
	"golang-bootcamp-1/pkg/utils"
	"time"
)

type IForgotPasswordUsecase interface {
	Create(dtoForgotPassword dto.ForgotPasswordRequest) (*entity.ForgotPassword, *response.ErrorResp)
	FindByCode(code string) (*entity.ForgotPassword, *response.ErrorResp)
	Update(dto dto.ForgotPasswordUpdateRequest) (*entity.ForgotPassword, *response.ErrorResp)
}

type forgotPasswordUsecase struct {
	repo   repo.IForgotPasswordRepo
	userUc userUc.IUserUseCase
	mail   mail.IMail
}

// Create usecase interact with forgot password repo
func (uc *forgotPasswordUsecase) Create(dtoForgotPassword dto.ForgotPasswordRequest) (*entity.ForgotPassword, *response.ErrorResp) {
	// Check user by email
	user, err := uc.userUc.FindByEmail(dtoForgotPassword.Email)
	if err != nil {
		return nil, err
	}

	// Set expired code
	// Store in UTC
	// Because the table store time without the timezone so it will always UTC
	expiryTime := time.
		Now().
		UTC().
		Add(24 * time.Hour)

	// Wrap the data, then save via repo
	data := entity.ForgotPassword{
		UserID:    &user.ID,
		Valid:     true,
		Code:      utils.RandomString(32),
		ExpiredAt: &expiryTime,
	}

	// Save data
	result, err := uc.repo.Create(data)
	if err != nil {
		return nil, err
	}

	// Send email
	emailData := dto.ForgotPasswordEmailBody{
		Subject:          "Forgot Password Code",
		Email:            user.Email,
		VerificationCode: result.Code,
	}

	go uc.mail.SendForgotPassword(user.Email, emailData)

	return result, nil
}

// Update implements IForgotPasswordUsecase.
func (uc *forgotPasswordUsecase) Update(dto dto.ForgotPasswordUpdateRequest) (*entity.ForgotPassword, *response.ErrorResp) {
	// Check code
	code, err := uc.repo.FindByCode(dto.Code)

	// Check invalidity of code
	// 1 - Error find code
	// 2 - Code not valid
	// 3 - Code expired
	if err != nil || !code.Valid || time.Now().After(*code.ExpiredAt) {
		return nil, &response.ErrorResp{
			Code:    404,
			Err:     errors.New("code is invalid"),
			Message: "Code is invalid or maybe expired!",
		}
	}

	// Search user
	user, err := uc.userUc.FindById(*code.UserID)
	if err != nil {
		return nil, err
	}

	// Create updated user data
	newUserData := userDto.UserUpdateRequest{
		Password: &dto.Password,
	}

	// Update user
	_, err = uc.userUc.Update(user.ID, newUserData)
	if err != nil {
		return nil, err
	}

	// Update forgot password
	code.Valid = false

	uc.repo.Update(*code)

	return code, nil
}

// FindByCode implements IForgotPasswordUsecase.
func (uc *forgotPasswordUsecase) FindByCode(code string) (*entity.ForgotPassword, *response.ErrorResp) {
	panic("unimplemented")
}

func NewForgotPasswordUsecase(repo repo.IForgotPasswordRepo, userUc userUc.IUserUseCase, mail mail.IMail) IForgotPasswordUsecase {
	return &forgotPasswordUsecase{
		repo,
		userUc,
		mail,
	}
}

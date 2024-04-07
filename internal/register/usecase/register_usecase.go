package register

import (
	registerDto "golang-bootcamp-1/internal/register/dto"
	userdto "golang-bootcamp-1/internal/user/dto"
	userUseCase "golang-bootcamp-1/internal/user/usecase"
	mail "golang-bootcamp-1/pkg/mail"
	"golang-bootcamp-1/pkg/response"
)

type IRegisterUsecase interface {
	Register(dto userdto.UserRequestBody) *response.ErrorResp
}

type registerUsecase struct {
	userUsecase userUseCase.IUserUseCase
	mailUc      mail.IMail
}

func NewRegisterUseCase(
	userUseCase userUseCase.IUserUseCase,
	mailUc mail.IMail,
) IRegisterUsecase {
	return &registerUsecase{
		userUsecase: userUseCase,
		mailUc:      mailUc,
	}
}

// Register implements IRegisterUsecase.
func (usecase *registerUsecase) Register(dto userdto.UserRequestBody) *response.ErrorResp {
	// Access other usecase
	// Access user usecase
	user, err := usecase.userUsecase.Create(dto)

	if err != nil {
		return err
	}

	// Send email confirm using sendgrid
	// Create data
	data := registerDto.EmailVerification{
		Subject:          "Verifikasi akun",
		Email:            user.Email,
		VerificationCode: user.CodeVerified,
	}

	// Create mail instance
	// Send email concurrently
	go usecase.mailUc.SendVerification(user.Email, data)

	return nil
}

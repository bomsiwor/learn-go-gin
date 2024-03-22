package register

import (
	userdto "golang-bootcamp-1/internal/user/dto"
	userUseCase "golang-bootcamp-1/internal/user/usecase"
	"golang-bootcamp-1/pkg/response"
)

type IRegisterUsecase interface {
	Register(dto userdto.UserRequestBody) *response.ErrorResp
}

type registerUsecase struct {
	userUsecase userUseCase.IUserUseCase
}

// Register implements IRegisterUsecase.
func (usecase *registerUsecase) Register(dto userdto.UserRequestBody) *response.ErrorResp {
	// Access other usecase
	// Access user usecase
	_, err := usecase.userUsecase.Create(dto)

	if err != nil {
		return err
	}

	// Send email confirm using sendgrid
	return nil

}

func NewRegisterUseCase(userUseCase userUseCase.IUserUseCase) IRegisterUsecase {
	return &registerUsecase{
		userUsecase: userUseCase,
	}
}

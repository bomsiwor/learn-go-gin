package register

import (
	registerUsecase "golang-bootcamp-1/internal/register/usecase"
	userDto "golang-bootcamp-1/internal/user/dto"
	"golang-bootcamp-1/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	uc registerUsecase.IRegisterUsecase
}

// Constructor for handler
func NewRegisterHandler(uc registerUsecase.IRegisterUsecase) *RegisterHandler {
	return &RegisterHandler{
		uc: uc,
	}
}

// Routing for register handler
func (handler *RegisterHandler) Router(r *gin.RouterGroup) {
	r.POST("register", handler.Register)
}

// Function for handling register
func (handler *RegisterHandler) Register(ctx *gin.Context) {
	// Validasi input

	// Bind dto with request
	var registerRequest userDto.UserRequestBody

	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			response.GenerateResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			),
		)
		ctx.Abort()
		return
	}

	// Start processing on usecase
	if err := handler.uc.Register(registerRequest); err != nil {
		ctx.JSON(
			err.Code,
			response.GenerateResponse(
				err.Code,
				err.Message,
				nil,
			),
		)
		ctx.Abort()
		return
	}

	// If success, send email later
	// Dont  forget to handle this

	// Return success message
	ctx.JSON(
		http.StatusCreated,
		response.GenerateResponse(
			http.StatusCreated,
			http.StatusText(http.StatusCreated),
			"Success, check your email",
		),
	)
}

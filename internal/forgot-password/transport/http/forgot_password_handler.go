package forgot_password

import (
	dto "golang-bootcamp-1/internal/forgot-password/dto"
	usecase "golang-bootcamp-1/internal/forgot-password/usecase"
	"golang-bootcamp-1/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordHandler struct {
	uc usecase.IForgotPasswordUsecase
}

func NewForgotPasswordHandler(usecase usecase.IForgotPasswordUsecase) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{
		uc: usecase,
	}
}

// Routing for forgot password handler
func (handler *ForgotPasswordHandler) Router(r *gin.RouterGroup) {
	forgotGroup := r.Group("forgot-password")
	forgotGroup.POST("create", handler.Create)
	forgotGroup.PATCH("update", handler.Update)
}

// Handler for create forgot password request
func (handler *ForgotPasswordHandler) Create(ctx *gin.Context) {
	// Validasi input

	// Bind data with request
	var forgotPasswordRequest dto.ForgotPasswordRequest

	if err := ctx.ShouldBindJSON(&forgotPasswordRequest); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			response.GenerateResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				err.Error(),
			),
		)

		// Abort context after throwing http response
		ctx.Abort()
		return
	}

	// Start process data via usecase
	_, err := handler.uc.Create(forgotPasswordRequest)
	if err != nil {
		ctx.JSON(
			err.Code,
			response.GenerateResponse(
				err.Code,
				err.Message,
				err.Err,
			),
		)

		ctx.Abort()
		return
	}

	// Return data
	ctx.JSON(
		http.StatusOK,
		response.GenerateResponse(
			http.StatusOK,
			"Success check email",
			true,
		),
	)
}

// Handler for updating password after user received the code
func (handler *ForgotPasswordHandler) Update(ctx *gin.Context) {
	var requestData dto.ForgotPasswordUpdateRequest

	// validating input
	// not implemented

	// bind request data to struct data
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			response.GenerateResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				err.Error(),
			),
		)

		ctx.Abort()
		return
	}

	// Start update data via usecase
	_, err := handler.uc.Update(requestData)
	if err != nil {
		ctx.JSON(
			err.Code,
			response.GenerateResponse(
				err.Code,
				http.StatusText(err.Code),
				err.Message,
			),
		)

		ctx.Abort()
		return
	}

	ctx.JSON(
		http.StatusOK,
		response.GenerateResponse(http.StatusOK, "Success updated password", true),
	)
}

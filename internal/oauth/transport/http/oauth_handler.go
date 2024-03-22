package oauth

import (
	dto "golang-bootcamp-1/internal/oauth/dto"
	usecase "golang-bootcamp-1/internal/oauth/usecase"
	"golang-bootcamp-1/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type oauthHandler struct {
	usecase usecase.IOauthUseCase
}

func NewOauthHandler(usecase usecase.IOauthUseCase) *oauthHandler {
	return &oauthHandler{
		usecase: usecase,
	}
}

func (handler *oauthHandler) Router(r *gin.RouterGroup) {
	groupRouter := r.Group("/api/v1")
	groupRouter.POST("/oauths", handler.Login)
}

func (handler *oauthHandler) Login(ctx *gin.Context) {
	var requestData dto.LoginRequest

	// Catch login request data
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			response.GenerateResponse(
				http.StatusInternalServerError,
				err.Error(),
				nil,
			),
		)
		ctx.Abort()
		return
	}

	// Call usecase
	// Pass the requested data to usecase
	loginResponse, err := handler.usecase.Login(requestData)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			err,
		)
		ctx.Abort()
		return
	}

	// Return success messsage if it OK!
	ctx.JSON(
		http.StatusOK,
		response.GenerateResponse(
			http.StatusOK,
			"Success login",
			loginResponse,
		),
	)
}

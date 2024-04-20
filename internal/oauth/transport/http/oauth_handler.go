package oauth

import (
	"fmt"
	"golang-bootcamp-1/internal/middleware"
	dto "golang-bootcamp-1/internal/oauth/dto"
	usecase "golang-bootcamp-1/internal/oauth/usecase"
	"golang-bootcamp-1/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type OauthHandler struct {
	usecase usecase.IOauthUseCase
}

func NewOauthHandler(usecase usecase.IOauthUseCase) *OauthHandler {
	return &OauthHandler{
		usecase: usecase,
	}
}

func (handler *OauthHandler) Router(r *gin.RouterGroup) {
	group := r.Group("oauth")
	group.POST("login", handler.Login)
	group.POST("refresh", handler.Refresh)

	group.Use(
		middleware.JwtTokenCheck,
		permissionMiddleware("create-admin"),
	).
		GET("me", handler.Me)
}

func (handler *OauthHandler) Login(ctx *gin.Context) {
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
			response.GenerateResponse(
				err.Code,
				err.Message,
				err.Err.Error(),
			),
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

func (handler *OauthHandler) Refresh(ctx *gin.Context) {
	var request dto.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
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

	// Refresh token via usecase
	result, err := handler.usecase.Refresh(request)
	if err != nil {
		ctx.JSON(
			err.Code,
			response.GenerateResponse(
				err.Code,
				err.Message,
				err.Err.Error(),
			),
		)
		ctx.Abort()
		return
	}

	// Return success messsage if it OK!
	ctx.JSON(
		http.StatusOK,
		response.GenerateResponse(
			http.StatusOK,
			"Success refresh token",
			result,
		),
	)
}

func (handler *OauthHandler) Me(ctx *gin.Context) {
	user, err := handler.usecase.Me(ctx.GetInt("user"), ctx.GetBool("isAdmin"))
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			response.GenerateResponse(
				http.StatusInternalServerError,
				err.Err.Error(),
				nil,
			),
		)
		ctx.Abort()
		return
	}

	ctx.JSON(
		http.StatusOK,
		response.GenerateResponse(
			http.StatusOK,
			"OK",
			user,
		),
	)
}

func (handler *OauthHandler) Logout(ctx *gin.Context) {

}

func permissionMiddleware(permission ...string) func(c *gin.Context) {
	questionMarksTmpl := make([]string, len(permission))
	for i := range permission {
		questionMarksTmpl[i] = "?"
	}

	questionMarks := strings.Join(questionMarksTmpl, ",")

	_ = fmt.Sprintf("select * from users where permission in (%s)", questionMarks)

	return func(c *gin.Context) {
		c.Next()
	}
}

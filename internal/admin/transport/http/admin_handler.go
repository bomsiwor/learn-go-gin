package admin

import (
	dto "golang-bootcamp-1/internal/admin/dto"
	uc "golang-bootcamp-1/internal/admin/usecase"
	"golang-bootcamp-1/internal/middleware"
	"golang-bootcamp-1/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	uc uc.IAdminUsecase
}

func NewAdminHandler(uc uc.IAdminUsecase) *AdminHandler {
	return &AdminHandler{
		uc: uc,
	}
}

func (h *AdminHandler) Router(r *gin.RouterGroup) {
	group := r.Group("admin").
		Use(middleware.IsAdmin)

	group.GET("", h.FindAll)
	group.GET(":id", h.FindByID)
	group.GET("total", h.TotalCountAdmin)
	group.POST("", h.Create)
	group.PUT(":id", h.Update)
	group.DELETE(":id", h.Delete)
}

func (h *AdminHandler) FindAll(ctx *gin.Context) {
	// parse query param
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	result := h.uc.FindAll(page, limit)

	ctx.JSON(
		http.StatusOK,
		response.GenerateResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			result,
		),
	)
}

func (h *AdminHandler) FindByID(ctx *gin.Context) {
	// Get ID
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Find user
	result, err := h.uc.FindByID(id)
	if err != nil {
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

	ctx.JSON(
		http.StatusOK,
		response.GenerateResponse(
			http.StatusOK,
			"User found",
			result,
		),
	)
}

func (h *AdminHandler) Create(ctx *gin.Context) {
	var request dto.AdminRequestBody

	// Validate Data

	if err := ctx.ShouldBindJSON(&request); err != nil {
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

	_, err := h.uc.Create(request)
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

	ctx.JSON(
		http.StatusCreated,
		response.GenerateResponse(
			http.StatusCreated,
			"Success create admin",
			true,
		),
	)
}

func (h *AdminHandler) Update(ctx *gin.Context) {
	var request dto.AdminRequestBody

	// Get ID
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Validate Data

	if err := ctx.ShouldBindJSON(&request); err != nil {
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

	_, err := h.uc.Update(id, request)
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

	ctx.JSON(
		http.StatusCreated,
		response.GenerateResponse(
			http.StatusCreated,
			"Success update admin",
			true,
		),
	)
}

func (h *AdminHandler) Delete(ctx *gin.Context) {
	// Get ID
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Delete data
	err := h.uc.Delete(id)
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

	ctx.JSON(
		http.StatusOK,
		response.GenerateResponse(
			http.StatusOK,
			"Admin deleted",
			true,
		),
	)
}

func (h *AdminHandler) TotalCountAdmin(ctx *gin.Context) {
	result, err := h.uc.TotalCountAdmin()
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

	ctx.JSON(
		http.StatusOK,
		response.GenerateResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			result,
		),
	)
}

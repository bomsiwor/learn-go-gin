package product_category

import (
	"golang-bootcamp-1/internal/middleware"
	dto "golang-bootcamp-1/internal/product-category/dto"
	usecase "golang-bootcamp-1/internal/product-category/usecase"
	"golang-bootcamp-1/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductCategoryHandler struct {
	uc usecase.IProductCategoryUsecase
}

func NewProductCategoryHandler(usecase usecase.IProductCategoryUsecase) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		uc: usecase,
	}
}

func (h *ProductCategoryHandler) Router(r *gin.RouterGroup) {
	group := r.Group("product-category").Use(middleware.JwtTokenCheck)

	group.GET("/", h.FindAll)
	group.GET(":id", h.FindById)
	group.POST("/", h.Store)
}

func (h *ProductCategoryHandler) FindAll(ctx *gin.Context) {
	// Get pagination query from URl
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

func (h *ProductCategoryHandler) FindById(ctx *gin.Context) {
	// Parse url Param
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			response.GenerateResponse(
				http.StatusInternalServerError,
				errParam.Error(),
				nil,
			),
		)

		ctx.Abort()
		return
	}

	result, err := h.uc.FindById(id)
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

	ctx.JSON(
		http.StatusOK,
		response.GenerateResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			result,
		),
	)
}

func (h *ProductCategoryHandler) Store(ctx *gin.Context) {
	// Parse form data to request struct
	var request dto.ProducteCategoryRequest

	if err := ctx.ShouldBind(&request); err != nil {
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

	// Assign user ID to creator
	userId := int64(ctx.GetInt("user"))
	request.CreatedBy = &userId

	result, errStore := h.uc.Create(request)
	if errStore != nil {
		ctx.JSON(
			errStore.Code,
			response.GenerateResponse(
				errStore.Code,
				errStore.Message,
				errStore.Err.Error(),
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

func (h *ProductCategoryHandler) Update(ctx *gin.Context) {

}

func (h *ProductCategoryHandler) Delete(ctx *gin.Context) {

}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"vandyahmad/newgotoko/auth"
	"vandyahmad/newgotoko/category"
	"vandyahmad/newgotoko/entity"
	"vandyahmad/newgotoko/helper"
)

type categoryHandler struct {
	categoryService category.Service
	authService     auth.Service
}

func NewCategoryHandler(authService auth.Service, categoryService category.Service) *categoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
		authService:     authService,
	}
}

func (h *categoryHandler) ListCategory(c *gin.Context) {

	// get query param
	limit := c.Query("limit")
	skip := c.Query("skip")
	if limit == "" {
		limit = strconv.Itoa(10)
	}
	if skip == "" {
		skip = strconv.Itoa(0)
	}
	lim, _ := strconv.Atoi(limit)
	skp, _ := strconv.Atoi(skip)
	result, err := h.categoryService.GetAll(lim, skp)
	if err != nil {
		response := helper.ApiResponse(false, "Category not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	meta := helper.Meta{
		Total: int(h.categoryService.CountAll()),
		Limit: lim,
		Skip:  skp,
	}
	responseCashier := category.CategoryListResponse{
		Category: result,
		Meta:     meta,
	}

	response := helper.ApiResponse(true, "Success", responseCashier)
	c.JSON(http.StatusOK, response)
	return

}

func (h *categoryHandler) DetailCategory(c *gin.Context) {
	id := c.Param("categoryId")
	idInt, _ := strconv.Atoi(id)
	result, err := h.categoryService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.ApiResponse(true, "Success", result)
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input category.InputCategory
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := entity.Response{
			Success: false,
			Message: "invalid request body",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	validate = validator.New()
	if err := validate.Struct(input); err != nil {
		response := helper.FormatErrorValidationCreate(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCashier, err := h.categoryService.RegisterCategory(&input)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse(true, "Success", newCashier)
	c.JSON(http.StatusOK, response)
	return

}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var input category.InputCategory
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := entity.Response{
			Success: false,
			Message: "invalid request body",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	validate = validator.New()
	if err := validate.Struct(input); err != nil {
		response := helper.FormatErrorValidationUpdate(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id := c.Param("categoryId")
	idInt, _ := strconv.Atoi(id)
	_, err = h.categoryService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	// maka update
	_, err = h.categoryService.UpdateCategory(idInt, &input)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiWithOutData(true, "Success")
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("categoryId")
	idInt, _ := strconv.Atoi(id)
	_, err := h.categoryService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	err = h.categoryService.Delete(idInt)
	response := helper.ApiWithOutData(true, "Success")
	c.JSON(http.StatusOK, response)
	return

}

package handler

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"vandyahmad/newgotoko/auth"
	"vandyahmad/newgotoko/cashier"
	"vandyahmad/newgotoko/entity"
	"vandyahmad/newgotoko/helper"

	"github.com/gin-gonic/gin"
)

type cashierHandler struct {
	cashierService cashier.Service
	authService    auth.Service
}

func NewCashierHandler(authService auth.Service, cashierService cashier.Service) *cashierHandler {
	return &cashierHandler{
		cashierService: cashierService,
		authService:    authService,
	}
}

var validate *validator.Validate

func (h *cashierHandler) ListCashier(c *gin.Context) {

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
	result, err := h.cashierService.GetAll(lim, skp)
	if err != nil {
		response := helper.ApiResponse(false, "Cashier not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	meta := helper.Meta{
		Total: int(h.cashierService.CountAll()),
		Limit: lim,
		Skip:  skp,
	}
	responseCashier := cashier.CashierListResponse{
		Cashiers: result,
		Meta:     meta,
	}

	response := helper.ApiResponse(true, "Success", responseCashier)
	c.JSON(http.StatusOK, response)
	return

}

func (h *cashierHandler) DetailCashier(c *gin.Context) {
	id := c.Param("cashierId")
	idInt, _ := strconv.Atoi(id)
	result, err := h.cashierService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.ApiResponse(true, "Success", result)
	c.JSON(http.StatusOK, response)
	return
}

func (h *cashierHandler) CreateCashier(c *gin.Context) {
	var input cashier.InputCashier
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

	newCashier, err := h.cashierService.RegisterCashier(&input)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse(true, "Success", newCashier)
	c.JSON(http.StatusOK, response)
	return

}

func (h *cashierHandler) UpdateCashier(c *gin.Context) {
	var input cashier.InputCashier
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

	id := c.Param("cashierId")
	idInt, _ := strconv.Atoi(id)
	_, err = h.cashierService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	// maka update
	_, err = h.cashierService.UpdateCashier(idInt, &input)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiWithOutData(true, "Success")
	c.JSON(http.StatusOK, response)
	return
}

func (h *cashierHandler) DeleteCashier(c *gin.Context) {
	id := c.Param("cashierId")
	idInt, _ := strconv.Atoi(id)
	_, err := h.cashierService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	err = h.cashierService.Delete(idInt)
	response := helper.ApiWithOutData(true, "Success")
	c.JSON(http.StatusOK, response)
	return

}

func (h *cashierHandler) GetPasscode(c *gin.Context) {
	id := c.Param("cashierId")
	idInt, _ := strconv.Atoi(id)
	casier, err := h.cashierService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	passcode := cashier.CashierPasscode{
		Passcode: casier.Passcode,
	}
	response := helper.ApiResponse(true, "Success", passcode)
	c.JSON(http.StatusOK, response)
	return
}

func (h *cashierHandler) LoginPasscode(c *gin.Context) {
	id := c.Param("cashierId")
	idInt, _ := strconv.Atoi(id)
	casier, err := h.cashierService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var input cashier.InputPasscode
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.FormatErrorValidationCreate(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Passcode != casier.Passcode {
		response := helper.ApiResponse(false, "Passcode Not Match", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	token, err := h.authService.GenerateToken(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	tokenResponse := cashier.TokenResponse{
		Token: token,
	}

	response := helper.ApiResponse(true, "Success", tokenResponse)
	c.JSON(http.StatusOK, response)
	return

}

func (h *cashierHandler) LogoutPasscode(c *gin.Context) {
	id := c.Param("cashierId")
	idInt, _ := strconv.Atoi(id)
	casier, err := h.cashierService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var input cashier.InputPasscode
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.FormatErrorValidationCreate(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Passcode != casier.Passcode {
		response := helper.ApiResponse(false, "Passcode Not Match", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	response := helper.ApiWithOutData(true, "Success")
	c.JSON(http.StatusOK, response)
	return

}

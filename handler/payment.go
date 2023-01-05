package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"vandyahmad/newgotoko/auth"
	"vandyahmad/newgotoko/entity"
	"vandyahmad/newgotoko/helper"
	"vandyahmad/newgotoko/payment"
)

type paymentHandler struct {
	paymentService payment.Service
}

func NewPaymentHandler(authService auth.Service, paymentService payment.Service) *paymentHandler {
	return &paymentHandler{
		paymentService: paymentService,
	}
}

func (h *paymentHandler) ListPayment(c *gin.Context) {

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
	result, err := h.paymentService.GetAll(lim, skp)
	if err != nil {
		response := helper.ApiResponse(false, "Payment not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	meta := helper.Meta{
		Total: int(h.paymentService.CountAll()),
		Limit: lim,
		Skip:  skp,
	}
	responsePayment := payment.PaymenttListResponse{
		Payments: result,
		Meta:     meta,
	}

	response := helper.ApiResponse(true, "Success", responsePayment)
	c.JSON(http.StatusOK, response)
	return

}

func (h *paymentHandler) DetailPayment(c *gin.Context) {
	id := c.Param("paymentId")
	idInt, _ := strconv.Atoi(id)
	result, err := h.paymentService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.ApiResponse(true, "Success", result)
	c.JSON(http.StatusOK, response)
	return
}

func (h *paymentHandler) CreatePayment(c *gin.Context) {
	var input payment.RequestPayment
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

	newCashier, err := h.paymentService.RegisterPayment(&input)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse(true, "Success", newCashier)
	c.JSON(http.StatusOK, response)
	return

}

func (h *paymentHandler) UpdatePayment(c *gin.Context) {
	var input payment.RequestPayment
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

	id := c.Param("paymentId")
	idInt, _ := strconv.Atoi(id)
	_, err = h.paymentService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	// maka update
	_, err = h.paymentService.UpdatePayment(idInt, &input)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiWithOutData(true, "Success")
	c.JSON(http.StatusOK, response)
	return
}

func (h *paymentHandler) DeletePayment(c *gin.Context) {
	id := c.Param("paymentId")
	idInt, _ := strconv.Atoi(id)
	_, err := h.paymentService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	err = h.paymentService.Delete(idInt)
	response := helper.ApiWithOutData(true, "Success")
	c.JSON(http.StatusOK, response)
	return

}

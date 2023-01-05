package router

import (
	"github.com/gin-gonic/gin"
	"vandyahmad/newgotoko/auth"
	"vandyahmad/newgotoko/config"
	"vandyahmad/newgotoko/handler"
	"vandyahmad/newgotoko/payment"
)

func PaymentRoute(app *gin.Engine) {
	authService := auth.NewServiceAuth()
	paymentRepository := payment.NewRepository(config.DB)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(authService, paymentService)

	app.GET("/payments", paymentHandler.ListPayment)
	app.GET("/payments/:paymentId", paymentHandler.DetailPayment)
	app.POST("/payments", paymentHandler.CreatePayment)
	app.PUT("/payments/:paymentId", paymentHandler.UpdatePayment)
	app.DELETE("/payments/:paymentId", paymentHandler.DeletePayment)

}

package router

import (
	"vandyahmad/gotoko/auth"
	"vandyahmad/gotoko/cashier"
	"vandyahmad/gotoko/config"
	"vandyahmad/gotoko/handler"

	"github.com/gin-gonic/gin"
)

func CashierRoute(app *gin.Engine) {
	authService := auth.NewServiceAuth()
	cashierRepository := cashier.NewRepository(config.DB)
	cashierService := cashier.NewService(cashierRepository)
	cashierHandler := handler.NewCashierHandler(authService, cashierService)

	app.GET("/cashiers", cashierHandler.ListCashier)
	app.GET("/cashiers/:cashierId", cashierHandler.DetailCashier)
	app.POST("/cashiers", cashierHandler.CreateCashier)
	app.PUT("/cashiers/:cashierId", cashierHandler.UpdateCashier)
	app.DELETE("/cashiers/:cashierId", cashierHandler.DeleteCashier)

	// login
	app.GET("/cashiers/:cashierId/passcode", cashierHandler.GetPasscode)
	app.POST("/cashiers/:cashierId/login", cashierHandler.LoginPasscode)
	app.POST("/cashiers/:cashierId/logout", cashierHandler.LogoutPasscode)
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vandyahmad/newgotoko/config"
	"vandyahmad/newgotoko/helper"
	rtr "vandyahmad/newgotoko/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Gotoko Test By Vandy Ahmad")
	config.InitDB()

	router := gin.Default()
	router.NoRoute(func(ctx *gin.Context) {
		response := helper.ApiResponse(false, "Error", nil)
		ctx.JSON(404, response)
	})
	rtr.CashierRoute(router)
	rtr.CategoryRoute(router)
	rtr.PaymentRoute(router)
	go func() {
		router.Run(":3030")
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("Proses Selesai dengan signal: %v\n", signal.String())
}

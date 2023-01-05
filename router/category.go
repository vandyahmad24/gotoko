package router

import (
	"github.com/gin-gonic/gin"
	"vandyahmad/newgotoko/auth"
	"vandyahmad/newgotoko/category"
	"vandyahmad/newgotoko/config"
	"vandyahmad/newgotoko/handler"
)

func CategoryRoute(app *gin.Engine) {
	authService := auth.NewServiceAuth()
	categoryRepository := category.NewRepository(config.DB)
	categoryService := category.NewService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(authService, categoryService)

	app.GET("/categories", categoryHandler.ListCategory)
	app.GET("/categories/:categoryId", categoryHandler.DetailCategory)
	app.POST("/categories", categoryHandler.CreateCategory)
	app.PUT("/categories/:categoryId", categoryHandler.UpdateCategory)
	app.DELETE("/categories/:categoryId", categoryHandler.DeleteCategory)

}

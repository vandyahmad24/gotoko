package middleware

import (
	"net/http"
	"strings"
	"vandyahmad/newgotoko/auth"
	"vandyahmad/newgotoko/cashier"
	"vandyahmad/newgotoko/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authMidd struct {
	authService   auth.Service
	casierService cashier.Service
}

func NewAuthMiddleware(casierService cashier.Service, authService auth.Service) *authMidd {
	return &authMidd{
		casierService: casierService,
		authService:   authService,
	}
}

func (a *authMidd) AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		response := helper.ApiResponse(false, "Unauthorized", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	var stringToken string
	arrayHeader := strings.Split(authHeader, " ")
	if len(arrayHeader) == 2 {
		stringToken = arrayHeader[1]
	}
	jwtToken, err := a.authService.ValidateToken(stringToken)
	if err != nil {
		response := helper.ApiResponse(false, "Unauthorized", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claim, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		response := helper.ApiResponse(false, "Unauthorized", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	userId := int(claim["id"].(float64))
	user, err := a.casierService.GetById(userId)
	if err != nil {
		response := helper.ApiResponse(false, "Unauthorized", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.Set("currentUser", user)

	// token, err :=
}

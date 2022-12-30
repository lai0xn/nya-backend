package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authcontroller *AuthController
}

func NewRouter(controller AuthController) *AuthRouter {
	return &AuthRouter{authcontroller: &controller}
}

func (ar *AuthRouter) Route(router *gin.Engine) {
	router.POST("/api/login", ar.authcontroller.Login)
	router.POST("/api/signup", ar.authcontroller.Signup)
}

package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jnxvi/nyalist/middlewares"
)

type UsersRouter struct {
	userscontroller UsersController
}

func NewRouter(controller *UsersController) *UsersRouter {
	return &UsersRouter{userscontroller: *controller}

}

func (r UsersRouter) Route(router *gin.Engine) {

	// protected routes
	router.Use(middlewares.LoginMiddleware())

	router.GET("/api/user/me", r.userscontroller.Profile)
	router.DELETE("/api/user/me/delete", r.userscontroller.Delete)
	router.PUT("/api/user/me/update", r.userscontroller.UpdateUser)

}

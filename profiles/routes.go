package profiles

import (
	"github.com/gin-gonic/gin"
	"github.com/jnxvi/nyalist/middlewares"
)

type ProfilesRouter struct {
	profileCotroller *ProfilesController
}

func NewRouter(controller ProfilesController) *ProfilesRouter {
	return &ProfilesRouter{profileCotroller: &controller}
}

func (ar *ProfilesRouter) Route(router *gin.Engine) {
	router.GET("/api/profiles/:username", ar.profileCotroller.GetUserProfile)

	//protected routes
	router.Use(middlewares.LoginMiddleware())
	router.PATCH("/api/profiles/me/update", ar.profileCotroller.UpdateProfile)
	router.POST("/api/profiles/add-to-list/:id", ar.profileCotroller.AddToWatchList)
	router.POST("/api/profiles/remove-from-list/:id", ar.profileCotroller.RemoveFromWatchList)

}

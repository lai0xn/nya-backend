package data

import (
	"github.com/gin-gonic/gin"
)

type DataRouter struct {
	dataontroller DataController
}

func NewRouter(controller *DataController) *DataRouter {
	return &DataRouter{dataontroller: *controller}

}

func (r DataRouter) Route(router *gin.Engine) {

	router.GET("/api/anime/search", r.dataontroller.SearchAnime)
	router.GET("/api/anime/watch", r.dataontroller.WatchLink)
	router.GET("/api/anime/download", r.dataontroller.DownloadLinks)
	router.GET("/api/anime/upcoming", r.dataontroller.UpcomingAnimes)
	router.GET("/api/anime/top", r.dataontroller.TopAnimes)

}

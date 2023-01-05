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

	router.GET("/api/search/anime", r.dataontroller.SearchAnime)
	router.GET("/api/watch/anime", r.dataontroller.WatchLink)
	router.GET("/api/download/anime", r.dataontroller.DownloadLinks)
	router.GET("/api/upcoming/anime", r.dataontroller.UpcomingAnimes)
	router.GET("/api/top/anime", r.dataontroller.TopAnimes)
	router.GET("/api/anime/:id", r.dataontroller.SearchByID)
	router.GET("/api/random/anime", r.dataontroller.RandomAnime)

}

package data

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DataController struct {
	DB *gorm.DB
}

var wrapper AnimeWrapper = AnimeWrapper{}

func NewController(DB *gorm.DB) *DataController {
	return &DataController{DB: DB}
}

func (DataController) SearchAnime(c *gin.Context) {
	anime_query := c.Query("query")

	animes, err := wrapper.SearchAnime(anime_query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Something went wrong")
	}
	c.JSON(http.StatusOK, gin.H{
		"data": animes,
	})
}

func (DataController) WatchLink(c *gin.Context) {
	anime_query := c.Query("query")
	episode := c.Query("episode")

	animes := wrapper.AnimeWatchLink(anime_query, episode)

	c.JSON(http.StatusOK, gin.H{
		"data": animes,
	})
}

func (DataController) DownloadLinks(c *gin.Context) {
	anime_query := c.Query("query")
	episode := c.Query("episode")

	animes := wrapper.DownloadLinks(anime_query, episode)

	c.JSON(http.StatusOK, gin.H{
		"download_links": animes,
	})
}

func (DataController) UpcomingAnimes(c *gin.Context) {
	animes, err := wrapper.GetUpcomingAnimes()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": animes,
	})
}
func (DataController) TopAnimes(c *gin.Context) {
	animes, err := wrapper.GetTopAnimes()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": animes,
	})
}

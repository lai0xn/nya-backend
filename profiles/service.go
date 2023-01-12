package profiles

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jnxvi/nyalist/data"
	"github.com/jnxvi/nyalist/models"

	"gorm.io/gorm"
)

type ProfilesController struct {
	DB *gorm.DB
}

func NewController(db *gorm.DB) *ProfilesController {
	return &ProfilesController{DB: db}
}

func (pc ProfilesController) GetUserProfile(ctx *gin.Context) {
	username := ctx.Param("username")
	var profile models.Profile
	var watchlist []models.Anime
	var user models.User
	pc.DB.Where("username = ?", username).Preload("Profile").Find(&user)
	if len(user.Email) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			"User does not exist")
		return
	}
	pc.DB.Where("profile_id = ?", user.ID).Find(&watchlist)
	pc.DB.Where("user_id = ?", user.ID).Preload("Watchlist").Find(&profile)
	ctx.JSON(http.StatusOK, gin.H{
		"username":  user.Username,
		"email":     user.Email,
		"bio":       user.Profile.Bio,
		"avatar":    user.Profile.ProfilePic,
		"watchlist": watchlist,
	})

}

func (pc ProfilesController) UpdateProfile(ctx *gin.Context) {
	var payload ProfilePaylaod
	var user models.User
	current_user, _ := ctx.Get("current_user")
	pc.DB.Where("id = ?", current_user.(models.User).ID).Find(&user)
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if len(payload.Username) > 0 {
		user.Username = payload.Username

	}
	if len(payload.Email) > 0 {
		user.Email = payload.Email

	}
	if len(payload.Bio) > 0 {
		user.Profile.Bio = payload.Bio

	}
	pc.DB.Save(&user)
	ctx.JSON(http.StatusOK, user)

}

func (pc ProfilesController) AddToWatchList(ctx *gin.Context) {
	id := ctx.Param("id")

	var profile models.Profile
	current_user, _ := ctx.Get("current_user")
	animeData := data.AnimeWrapper{}
	anime, err := animeData.SearchAnimeByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	pc.DB.Where("id = ?", current_user.(models.User).ID).
		Find(&profile)
	fmt.Println(profile)
	pc.DB.Create(&models.Anime{MalID: anime.Data.MalId,
		Title:     anime.Data.Name,
		Image:     anime.Data.Images.Image.ImageURL,
		ProfileID: int(profile.ID)})

	ctx.JSON(http.StatusOK, "Anime Added to watchlist")

}

func (pc *ProfilesController) RemoveFromWatchList(ctx *gin.Context) {

	id := ctx.Param("id")

	current_user, _ := ctx.Get("current_user")
	var anime models.Anime

	pc.DB.Where("id = ? AND mal_id = ?", current_user.(models.User).ID, id).Find(&anime)

	pc.DB.Delete(&anime)

	ctx.JSON(http.StatusOK, "Anime Added to watchlist")

}

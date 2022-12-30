package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jnxvi/nyalist/auth"
	"gorm.io/gorm"
)

var payload UserPayload

type UsersController struct {
	DB *gorm.DB
}

func NewController(db *gorm.DB) *UsersController {
	return &UsersController{DB: db}

}

func (uc *UsersController) Profile(ctx *gin.Context) {
	current_user, _ := ctx.Get("current_user")

	ctx.JSON(http.StatusOK, gin.H{
		"username": current_user.(auth.User).Username,
		"email":    current_user.(auth.User).Email,
	})
}

func (uc *UsersController) Delete(ctx *gin.Context) {
	current_user, _ := ctx.Get("current_user")
	uc.DB.Delete(&auth.User{}, current_user.(auth.User).ID)
	ctx.JSON(http.StatusOK, "Your Account is Deleted")

}

func (uc *UsersController) UpdateUser(ctx *gin.Context) {
	current_user, _ := ctx.Get("current_user")
	var user auth.User
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	uc.DB.First(&user, "id = ?", current_user.(auth.User).ID)
	user.Email = payload.Email
	user.Username = payload.Username
	uc.DB.Save(&user)
	ctx.JSON(http.StatusOK, "User Updated")

}

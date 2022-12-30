package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jnxvi/nyalist/auth"
	"github.com/jnxvi/nyalist/database"
)

var authorizationHeaderKey = "authorization"
var authorizationHeaderType = "token"
var token string

func LoginMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user auth.User
		authorizationheader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationheader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Authorization token not provided")
			return
		}

		fields := strings.Fields(authorizationheader)

		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Invalid authorization format")
			return

		}
		token = fields[1]
		if fields[0] != authorizationHeaderType {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Invalid authorization type")
			return

		}
		database.DB.Find(&user, "auth_token = ?", token)
		if user.Email == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Invalid Token")
			return

		}
		ctx.Set("current_user", user)
		ctx.Next()
	}
}

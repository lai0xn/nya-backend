package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jnxvi/nyalist/auth"
	"github.com/jnxvi/nyalist/data"
	"github.com/jnxvi/nyalist/database"
	"github.com/jnxvi/nyalist/models"
	"github.com/jnxvi/nyalist/profiles"
	"github.com/jnxvi/nyalist/users"
)

var (
	authrouter         *auth.AuthRouter
	authcontroller     *auth.AuthController
	userscontroller    *users.UsersController
	usersrouter        *users.UsersRouter
	datacontroller     *data.DataController
	datarouter         *data.DataRouter
	profilescontroller *profiles.ProfilesController
	profilesrouter     *profiles.ProfilesRouter

	r *gin.Engine = gin.Default()
)

func main() {
	database.Connect()
	r.Static("/assets", "./static")
	authcontroller = auth.NewController(database.DB)
	authrouter = auth.NewRouter(*authcontroller)
	userscontroller = users.NewController(database.DB)
	usersrouter = users.NewRouter(userscontroller)
	datacontroller = data.NewController(database.DB)
	datarouter = data.NewRouter(datacontroller)
	profilescontroller = profiles.NewController(database.DB)
	profilesrouter = profiles.NewRouter(*profilescontroller)
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Profile{})
	database.DB.AutoMigrate(&models.Anime{})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:3000"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	datarouter.Route(r)

	authrouter.Route(r)
	profilesrouter.Route(r)

	usersrouter.Route(r)

	r.Run(":5050")
}

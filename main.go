package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jnxvi/nyalist/auth"
	"github.com/jnxvi/nyalist/data"
	"github.com/jnxvi/nyalist/database"
	"github.com/jnxvi/nyalist/users"
)

var (
	authrouter      *auth.AuthRouter
	authcontroller  *auth.AuthController
	userscontroller *users.UsersController
	usersrouter     *users.UsersRouter
	datacontroller  *data.DataController
	datarouter      *data.DataRouter

	r *gin.Engine = gin.Default()
)

func main() {
	database.Connect()
	authcontroller = auth.NewController(database.DB)
	authrouter = auth.NewRouter(*authcontroller)
	userscontroller = users.NewController(database.DB)
	usersrouter = users.NewRouter(userscontroller)
	datacontroller = data.NewController(database.DB)
	datarouter = data.NewRouter(datacontroller)
	database.DB.AutoMigrate(&auth.User{})
	datarouter.Route(r)

	authrouter.Route(r)
	usersrouter.Route(r)

	r.Run(":5050")
}

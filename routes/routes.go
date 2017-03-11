package routes

import (
	controllers "github.com/didiyudha/go-mvc-pattern/controllers"
	"github.com/julienschmidt/httprouter"
)

var uc *controllers.UserController

func init() {
	uc = controllers.NewUserController()
}

// NewRouter will return pointer to a router
func NewRouter() (router *httprouter.Router) {
	router = httprouter.New()
	router.GET("/", controllers.Index)
	router.GET("/users", uc.GetAllUsers)
	router.GET("/registers", controllers.Register)
	router.POST("/createUser", uc.CreateUser)
	return router
}

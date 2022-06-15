package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zepyrshut/gin-gonic-starter/internal/config"
	"github.com/zepyrshut/gin-gonic-starter/internal/handlers"
	"github.com/zepyrshut/gin-gonic-starter/internal/middleware"
	"github.com/zepyrshut/gin-gonic-starter/internal/util"
)

var app *config.Application

func NewRoutes(a *config.Application) {
	app = a
}

func Routes() *gin.Engine {

	router := gin.Default()

	// CORS and CSRF protection
	router.Use(cors.Default())
	router.Use(middleware.Localize())
	router.Use(middleware.Sessions("session"))
	router.Use(middleware.CORSMiddleware())

	// Status and test
	router.GET("/status", handlers.Repo.GetStatusHandler)
	router.GET("/localize", util.GetLocalize)

	// User
	router.GET("/user/:id", handlers.Repo.GetOneUser)

	//Testing sessions and CSRF protection
	router.GET("/incr", util.Increment)
	router.GET("/protected", util.GetToken)
	router.POST("/protected", util.PostToken)

	return router

}

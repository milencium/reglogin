package router

import (
	"github.com/milencium/reglogin/controller"
	"github.com/milencium/reglogin/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter setup routing here
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORSMiddleware())

	// routes
	router.GET("/ping", controller.Pong)
	router.POST("/register", controller.Create)
	router.POST("/login", controller.Login)
	router.GET("/session", controller.Session)
	router.POST("/createReset", controller.InitiatePasswordReset)
	router.POST("/resetPassword",controller.ResetPassword)
	return router
}

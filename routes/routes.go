package routes

import (
	"be-tesis/controllers"
	"be-tesis/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/register", func(c *gin.Context) {
		controllers.RegisterUser(c.Writer, c.Request)
	})
	router.POST("/login", func(c *gin.Context) {
		controllers.LoginUser(c.Writer, c.Request)
	})
	router.POST("/logout", func(c *gin.Context) {
		controllers.LogoutUser(c.Writer, c.Request)
	})
	router.GET("/users", middleware.AuthMiddleware(), func(c *gin.Context) {
		controllers.GetAllUsers(c.Writer, c.Request)
	})
}

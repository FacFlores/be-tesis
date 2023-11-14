// main.go
package main

import (
	"be-tesis/db"
	"be-tesis/models"
	"be-tesis/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB Connection
	db.DBConnection()
	router := gin.Default()
	//Creation of tables for database
	db.DB.AutoMigrate(models.User{})

	//Routes
	routes.UserRoutes(router)
	//port number
	router.Run(":8080")
}

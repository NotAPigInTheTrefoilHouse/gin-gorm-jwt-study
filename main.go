package main

import (
	"github.com/NotAPigInTheTrefoilHouse/gin-gorm-jwt-study/controllers"
	"github.com/NotAPigInTheTrefoilHouse/gin-gorm-jwt-study/initializers"
	"github.com/NotAPigInTheTrefoilHouse/gin-gorm-jwt-study/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}

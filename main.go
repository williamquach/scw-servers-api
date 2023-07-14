package main

import (
	"os"
	"servers_api/controllers"
	"servers_api/initializers"
	"servers_api/repositories"

	docs "servers_api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func initialize() *gorm.DB {
	initializers.LoadEnvironmentVariables()
	db := initializers.ConnectDatabase()

	return db
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func initRoutes(router *gin.Engine, db *gorm.DB) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", HealthCheckHandler)

	api := router.Group("/api/v1")

	serverRepository := repositories.GORMServerRepository{
		DB: db,
	}
	controllers.InitiateServersRoutes(api, serverRepository)
}

func initSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Title = "SCW Homework - Servers API"
	docs.SwaggerInfo.Description = "This API is used to manage servers"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/openapi/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func launchServer(router *gin.Engine) {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}

func main() {
	router := initializers.SetUpRouter()

	configApp := cors.DefaultConfig()
	configApp.AllowAllOrigins = true
	configApp.AllowHeaders = []string{"*"}
	configApp.AllowMethods = []string{"*"}
	configApp.MaxAge = 0
	router.Use(cors.New(configApp))

	db := initialize()
	initRoutes(router, db)
	initSwagger(router)
	launchServer(router)
}

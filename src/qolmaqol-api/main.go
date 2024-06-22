package main

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/qolmaqol-api/initializers"
	"log"
	"net/http"
)

var (
	server        *gin.Engine
	storageClient *storage.Client

	//AuthController      controllers.AuthController
	//AuthRouteController routes.AuthRouteController

	//StorageController      controllers.StorageController
	//StorageRouteController routes.StorageRouteController

	//UserController      controllers.UserController
	//UserRouteController routes.UserRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	initializers.ConnectDB(&config)

	storageClient, err = initializers.ConnectStorage()
	if err != nil {
		log.Fatal("? Could not connect Firebase Storage", err)
	}

	//AuthController = controllers.NewAuthController(initializers.DB)
	//AuthRouteController = routes.NewAuthRouteController(AuthController)

	//StorageController = controllers.NewStorageController(initializers.DB, storageClient)
	//StorageRouteController = routes.NewStorageRouteController(StorageController)

	//UserController = controllers.NewUserController(initializers.DB)
	//UserRouteController = routes.NewUserRouteController(UserController)

	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	server = gin.Default()
	server.Use(corsMiddleware())
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	router := server.Group("/api/v1")
	// PING method to check service status
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "pong"})
	})

	//AuthRouteController.AuthRoute(router)
	//StorageRouteController.StorageRoute(router)
	//UserRouteController.UserRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Q-Auth-Token, Q-Device-Id")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

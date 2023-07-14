package controllers

import (
	"servers_api/repositories"
	"servers_api/usecases"

	"github.com/gin-gonic/gin"
)

func InitiateServersRoutes(router *gin.RouterGroup, serverRepository repositories.ServerRepository) {
	createServerUseCase := usecases.CreateServerUseCase{
		ServerRepository: serverRepository,
	}
	findAllServersUseCase := usecases.FindAllServersUseCase{
		ServerRepository: serverRepository,
	}
	findServerByIDUseCase := usecases.FindServerByIDUseCase{
		ServerRepository: serverRepository,
	}

	serverController := ServerController{
		CreateServerUseCase:   createServerUseCase,
		FindAllServersUseCase: findAllServersUseCase,
		FindServerByIDUseCase: findServerByIDUseCase,
	}

	router.POST("/servers", serverController.CreateServer)
	router.GET("/servers", serverController.FindAllServers)
	router.GET("/servers/:id", serverController.FindServerByID)
}

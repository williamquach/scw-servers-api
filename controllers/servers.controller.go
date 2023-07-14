package controllers

import (
	"servers_api/controllers/requests"
	"servers_api/controllers/responses"
	"servers_api/usecases"
	"servers_api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	CreateServerUseCase   usecases.CreateServerUseCase
	FindAllServersUseCase usecases.FindAllServersUseCase
	FindServerByIDUseCase usecases.FindServerByIDUseCase
}

// CreateServer godoc
// @Summary Create a server
// @Description Create a server
// @Tags Servers
// @Accept  json
// @Produce  json
// @Param server body requests.CreateServerRequest true "Server"
// @Success 201 {object} responses.ServerResponse
// @Failure 400 {object} responses.APIError
// @Failure 500 {object} responses.APIError
// @Router /servers [post]
// CreateServer creates a server
func (s *ServerController) CreateServer(c *gin.Context) {
	var createServerRequest requests.CreateServerRequest

	if err := c.ShouldBindJSON(&createServerRequest); err != nil {
		responses.NewAPIError().
			WithMessage("Error parsing JSON").
			WithCode("invalid_request").
			WithStatus(400).
			WithContext(map[string]interface{}{
				"body": c.Request.Body,
			}).
			WithError(err).
			ProduceAPIError(c)
		return
	}

	createServer := createServerRequest.ToCommand()

	createdServer, err := s.CreateServerUseCase.Execute(createServer)
	if err != nil {
		responses.NewAPIError().
			WithMessage("Error creating server").
			WithCode("server_error").
			WithStatus(500).
			WithContext(map[string]interface{}{
				"createServer": createServer,
			}).
			WithError(err).
			ProduceAPIError(c)
		return
	}

	utils.ProduceCreatedAPIResponse(c, createdServer.ToResponse())
}

// FindAllServers godoc
// @Summary Find all servers
// @Description Find all servers
// @Tags Servers
// @Accept  json
// @Produce  json
// @Success 200 {array} responses.ServerResponse
// @Failure 500 {object} responses.APIError
// @Router /servers [get]
// FindAllServers Finds all servers
func (s *ServerController) FindAllServers(c *gin.Context) {
	servers, err := s.FindAllServersUseCase.Execute()
	if err != nil {
		responses.NewAPIError().
			WithMessage("Error getting all servers").
			WithCode("server_error").
			WithStatus(500).
			WithError(err).
			ProduceAPIError(c)
		return
	}

	var serversResponse []responses.ServerResponse
	for _, server := range servers {
		serversResponse = append(serversResponse, server.ToResponse())
	}
	utils.ProduceSuccessAPIResponse(c, serversResponse)
}

// FindServerByID godoc
// @Summary Find server by ID
// @Description Find server by ID
// @Tags Servers
// @Accept  json
// @Produce  json
// @Param id path string true "Server ID"
// @Success 200 {object} responses.ServerResponse
// @Failure 404 {object} responses.APIError
// @Failure 500 {object} responses.APIError
// @Router /servers/{id} [get]
// FindServerByID finds a server by ID
func (s *ServerController) FindServerByID(c *gin.Context) {
	serverID := c.Param("id")
	if serverID == "" {
		responses.NewAPIError().
			WithMessage("Server ID is required").
			WithCode("invalid_request").
			WithStatus(400).
			ProduceAPIError(c)
		return
	}

	convertedServerID, err := strconv.ParseUint(serverID, 10, 64)
	if err != nil {
		responses.NewAPIError().
			WithMessage("Error parsing server ID").
			WithCode("invalid_request").
			WithStatus(400).
			WithContext(map[string]interface{}{
				"serverId": serverID,
			}).
			WithError(err).
			ProduceAPIError(c)
		return
	}

	server, err := s.FindServerByIDUseCase.Execute(uint(convertedServerID))
	if err != nil {
		responses.NewAPIError().
			WithMessage("Error getting server").
			WithCode("server_error").
			WithStatus(500).
			WithContext(map[string]interface{}{
				"serverId": serverID,
			}).
			WithError(err).
			ProduceAPIError(c)
		return
	}

	if server == nil {
		responses.NewAPIError().
			WithMessage("Server not found").
			WithCode("not_found").
			WithStatus(404).
			WithContext(map[string]interface{}{
				"serverId": serverID,
			}).
			ProduceAPIError(c)
		return
	}

	utils.ProduceSuccessAPIResponse(c, server.ToResponse())
}

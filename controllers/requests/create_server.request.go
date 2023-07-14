package requests

import "servers_api/domain"

// CreateServerRequest represents the request to create a server
// swagger:model CreateServerRequest
type CreateServerRequest struct {
	// Name of the server
	// example: ESGI
	Name string `json:"name" binding:"required"`
	// Type of the server
	// example: SMALL | MEDIUM | LARGE
	Type domain.ServerType `json:"type" binding:"required,oneof=SMALL MEDIUM LARGE"`
	// Status of the server
	// example: STARTED | RUNNING | STOPPING | STOPPED
	Status domain.ServerStatus `json:"status" binding:"required,oneof=STARTING RUNNING STOPPING STOPPED"`
}

func (request CreateServerRequest) ToCommand() domain.CreateServer {
	return domain.CreateServer{
		Name:   request.Name,
		Type:   request.Type,
		Status: request.Status,
	}
}

package domain

import "servers_api/controllers/responses"

type Server struct {
	ID     uint         `json:"id"`
	Name   string       `json:"name"`
	Type   ServerType   `json:"type"`
	Status ServerStatus `json:"status"`
}

func (s Server) ToResponse() responses.ServerResponse {
	return responses.ServerResponse{
		ID:     s.ID,
		Name:   s.Name,
		Type:   string(s.Type),
		Status: string(s.Status),
	}
}

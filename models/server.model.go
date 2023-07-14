package models

import (
	"servers_api/domain"

	"gorm.io/gorm"
)

type ServerModel struct {
	gorm.Model
	Name   string              `json:"name"`
	Type   domain.ServerType   `json:"type"`
	Status domain.ServerStatus `json:"status"`
}

func (server ServerModel) TableName() string {
	return "servers"
}

func (server ServerModel) ToDomain() *domain.Server {
	return &domain.Server{
		ID:     server.ID,
		Name:   server.Name,
		Type:   server.Type,
		Status: server.Status,
	}
}

package repositories

import (
	"servers_api/domain"
	"servers_api/models"
)

type InMemoryServerRepository struct {
	Servers []models.ServerModel
}

func (repository *InMemoryServerRepository) Create(createServer domain.CreateServer) (*domain.Server, error) {
	server := models.ServerModel{
		Name:   createServer.Name,
		Type:   createServer.Type,
		Status: createServer.Status,
	}

	repository.Servers = append(repository.Servers, server)

	return server.ToDomain(), nil
}

func (repository InMemoryServerRepository) FindAll() ([]*domain.Server, error) {
	foundServers := repository.Servers

	var servers []*domain.Server
	for _, server := range foundServers {
		servers = append(servers, server.ToDomain())
	}

	return servers, nil
}

func (repository InMemoryServerRepository) FindByID(id uint) (*domain.Server, error) {
	var server *models.ServerModel

	for _, s := range repository.Servers {
		if s.ID == id {
			server = &s
			break
		}
	}

	if server == nil {
		return nil, nil
	}

	return server.ToDomain(), nil
}

package repositories

import (
	"servers_api/domain"
	"servers_api/models"

	"gorm.io/gorm"
)

type GORMServerRepository struct {
	DB *gorm.DB
}

func (repository GORMServerRepository) Create(createServer domain.CreateServer) (*domain.Server, error) {
	server := models.ServerModel{
		Name:   createServer.Name,
		Type:   createServer.Type,
		Status: createServer.Status,
	}

	result := repository.DB.Create(&server)

	if result.Error != nil {
		return nil, result.Error
	}

	return server.ToDomain(), nil
}

func (repository GORMServerRepository) FindAll() ([]*domain.Server, error) {
	var foundServers []models.ServerModel

	result := repository.DB.Find(&foundServers)

	if result.Error != nil {
		return nil, result.Error
	}

	servers := make([]*domain.Server, len(foundServers))
	for index, server := range foundServers {
		servers[index] = server.ToDomain()
	}

	return servers, nil
}

func (repository GORMServerRepository) FindByID(id uint) (*domain.Server, error) {
	var server models.ServerModel

	result := repository.DB.First(&server, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return server.ToDomain(), nil
}

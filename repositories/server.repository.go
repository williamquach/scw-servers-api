package repositories

import (
	"servers_api/domain"
)

type ServerRepository interface {
	Create(createServer domain.CreateServer) (*domain.Server, error)
	FindAll() ([]domain.Server, error)
	FindByID(id uint) (*domain.Server, error)
}

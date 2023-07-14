package usecases

import (
	"fmt"
	"servers_api/domain"
	"servers_api/repositories"
)

type FindServerByIDUseCase struct {
	ServerRepository repositories.ServerRepository
}

func (useCase FindServerByIDUseCase) Execute(serverId uint) (*domain.Server, error) {
	foundServer, err := useCase.ServerRepository.FindByID(serverId)
	if err != nil {
		return nil, fmt.Errorf("error while getting all servers: %w", err)
	}

	return foundServer, nil
}

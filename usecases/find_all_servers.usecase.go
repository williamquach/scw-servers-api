package usecases

import (
	"fmt"
	"servers_api/domain"
	"servers_api/repositories"
)

type FindAllServersUseCase struct {
	ServerRepository repositories.ServerRepository
}

func (useCase FindAllServersUseCase) Execute() ([]domain.Server, error) {
	foundServers, err := useCase.ServerRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error while getting all servers: %w", err)
	}

	return foundServers, nil
}

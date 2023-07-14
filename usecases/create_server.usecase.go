package usecases

import (
	"fmt"
	"servers_api/domain"
	"servers_api/repositories"
)

type CreateServerUseCase struct {
	ServerRepository repositories.ServerRepository
}

func (useCase CreateServerUseCase) Execute(createServer domain.CreateServer) (*domain.Server, error) {
	createdServer, err := useCase.ServerRepository.Create(createServer)
	if err != nil {
		return nil, fmt.Errorf("error while creating server: %w", err)
	}

	return createdServer, nil
}

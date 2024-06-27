package playerusecase

import playerrepository "microshop/modules/player/playerRepository"

type (
	PlayerUsecaseService interface{}

	playerUsecase struct {
		playerRepository playerrepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase(playerRepository playerrepository.PlayerRepositoryService) PlayerUsecaseService {
	return &playerUsecase{playerRepository: playerRepository}
}

package playerhandler

import (
	"microshop/config"
	playerusecase "microshop/modules/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface{}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(cfg *config.Config, playerUsecase playerusecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{playerUsecase: playerUsecase}
}

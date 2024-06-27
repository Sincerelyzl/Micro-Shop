package playerhandler

import playerusecase "microshop/modules/player/playerUsecase"

type (
	playerGrpcHandlerService struct {
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(playerUsecase playerusecase.PlayerUsecaseService) playerGrpcHandlerService {
	return playerGrpcHandlerService{playerUsecase}
}

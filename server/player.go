package server

import (
	playerhandler "microshop/modules/player/playerHandler"
	playerrepository "microshop/modules/player/playerRepository"
	playerusecase "microshop/modules/player/playerUsecase"
)

func (s *server) playerService() {
	repo := playerrepository.NewPlayerRepository(s.db)
	usecase := playerusecase.NewPlayerUsecase(repo)
	httpHandler := playerhandler.NewPlayerHttpHandler(s.cfg, usecase)
	grapcHandlder := playerhandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerhandler.NewPlayerQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grapcHandlder
	_ = queueHandler

	player := s.app.Group("/player_v1")

	//Health check
	_ = player
}

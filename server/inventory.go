package server

import (
	inventoryhandler "microshop/modules/inventory/inventoryHandler"
	inventoryrepository "microshop/modules/inventory/inventoryRepository"
	inventoryusecase "microshop/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryrepository.NewInventoryRepository(s.db)
	usecase := inventoryusecase.NewInventoryUsecase(repo)
	httpHandler := inventoryhandler.NewInventoryHttpHandler(s.cfg, usecase)
	grapcHandlder := inventoryhandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryhandler.NewInventoryQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grapcHandlder
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	//Health check
	_ = inventory
}

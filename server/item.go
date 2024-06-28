package server

import (
	itemhandler "microshop/modules/item/itemHandler"
	itemrepository "microshop/modules/item/itemRepository"
	itemusecase "microshop/modules/item/itemUsecase"
)

func (s *server) itemService() {
	repo := itemrepository.NewItemRepository(s.db)
	usecase := itemusecase.NewItemUsecase(repo)
	httpHandler := itemhandler.NewItemHttpHandler(s.cfg, usecase)
	grapcHandlder := itemhandler.NewItemGrpcHandler(usecase)

	_ = httpHandler
	_ = grapcHandlder

	item := s.app.Group("/item_v1")

	//Health check
	_ = item
}

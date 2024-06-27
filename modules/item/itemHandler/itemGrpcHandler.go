package itemhandler

import (
	"microshop/config"
	itemusecase "microshop/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemUsecase itemusecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(cfg *config.Config, itemUsecase itemusecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUsecase: itemUsecase}
}

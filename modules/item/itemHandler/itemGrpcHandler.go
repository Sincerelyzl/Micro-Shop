package itemhandler

import (
	itemusecase "microshop/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemUsecase itemusecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(itemUsecase itemusecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUsecase: itemUsecase}
}

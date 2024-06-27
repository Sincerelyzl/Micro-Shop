package itemhandler

import (
	"microshop/config"
	itemusecase "microshop/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface{}

	itemHttpHandler struct {
		cfg         *config.Config
		itemUsecase itemusecase.ItemUsecaseService
	}
)

func NewItemHandler(cfg *config.Config, itemUsecase itemusecase.ItemUsecaseService) ItemHttpHandlerService {
	return &itemHttpHandler{
		cfg:         cfg,
		itemUsecase: itemUsecase}
}

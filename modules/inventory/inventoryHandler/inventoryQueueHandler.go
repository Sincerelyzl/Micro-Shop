package inventoryhandler

import (
	"microshop/config"
	inventoryusecase "microshop/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface{}

	inventoryQueueHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler(cfg *config.Config, inventoryUsecase inventoryusecase.InventoryUsecaseService) InventoryQueueHandlerService {
	return &inventoryQueueHandler{
		cfg:              cfg,
		inventoryUsecase: inventoryUsecase}
}

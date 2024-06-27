package inventoryhandler

import (
	inventoryusecase "microshop/modules/inventory/inventoryUsecase"
)

type (
	InventoryGrpcHandlerService interface{}

	inventoryGrpcHandler struct {
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryusecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{
		inventoryUsecase: inventoryUsecase}
}

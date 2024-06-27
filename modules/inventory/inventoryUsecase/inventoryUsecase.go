package inventoryusecase

import inventoryrepository "microshop/modules/inventory/inventoryRepository"

type (
	InventoryUsecaseService interface{}

	inventoryUsecase struct {
		inventoryRepository inventoryrepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase(inventoryRepository inventoryrepository.InventoryRepositoryService) InventoryUsecaseService {
	return &inventoryUsecase{inventoryRepository: inventoryRepository}
}

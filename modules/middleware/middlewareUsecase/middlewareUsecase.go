package middlewareUsecase

import "microshop/modules/middleware/middlewareRepository"

type (
	MiddleweareUsecaseService interface{}

	middlewareUsecase struct {
		midwareRepository middlewareRepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepositoryService) MiddleweareUsecaseService {
	return &middlewareUsecase{middlewareRepository}
}

package middlewarehandler

import (
	"microshop/config"
	"microshop/modules/middleware/middlewareUsecase"
)

type (
	MiddleareHandlerService interface{}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase middlewareUsecase.MiddleweareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddleweareUsecaseService) MiddleareHandlerService {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}

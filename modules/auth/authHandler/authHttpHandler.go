package authHandler

import (
	"microshop/config"
	"microshop/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface{}

	AuthHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &AuthHttpHandler{cfg, authUsecase}
}

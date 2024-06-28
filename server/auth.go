package server

import (
	"microshop/modules/auth/authHandler"
	"microshop/modules/auth/authRepository"
	"microshop/modules/auth/authUsecase"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewwAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grapcHandlder := authHandler.NewAuthGrpcHandler(usecase)

	_ = httpHandler
	_ = grapcHandlder

	auth := s.app.Group("/auth_v1")

	//Health check
	_ = auth
}

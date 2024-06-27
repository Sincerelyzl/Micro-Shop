package authUsecase

import "microshop/modules/auth/authRepository"

type (
	AuthUsecaseService interface{}

	authUsecase struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewwAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}

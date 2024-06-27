package paymenthandler

import (
	"microshop/config"
	paymentusecase "microshop/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentusecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{cfg: cfg, paymentUsecase: paymentUsecase}
}

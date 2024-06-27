package paymenthandler

import (
	"microshop/config"
	paymentusecase "microshop/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		cfg            *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentusecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{cfg: cfg, paymentUsecase: paymentUsecase}
}

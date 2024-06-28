package server

import (
	paymenthandler "microshop/modules/payment/paymentHandler"
	paymentrepository "microshop/modules/payment/paymentRepository"
	paymentusecase "microshop/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentrepository.NewPaymentRepository(s.db)
	usecase := paymentusecase.NewPaymentUsecase(repo)
	httpHandler := paymenthandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymenthandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	//Health check
	_ = payment
}

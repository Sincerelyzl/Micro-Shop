package server

import (
	"context"
	"log"
	"microshop/config"
	middlewarehandler "microshop/modules/middleware/middlewareHandler"
	"microshop/modules/middleware/middlewareRepository"
	"microshop/modules/middleware/middlewareUsecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		app        *echo.Echo
		db         *mongo.Client
		cfg        *config.Config
		middleware middlewarehandler.MiddlewareHandlerService
	}
)

func newMiddleware(cfg *config.Config) middlewarehandler.MiddlewareHandlerService {
	repo := middlewareRepository.NewMiddlewareRepository()
	usecase := middlewareUsecase.NewMiddlewareUsecase(repo)
	return middlewarehandler.NewMiddlewareHandler(cfg, usecase)
}

func (s *server) gracefulShutdown(pctx context.Context, quit <-chan os.Signal) {
	log.Printf("Start service : %s", s.cfg.App.Name)
	<-quit
	log.Printf("Shutting down service : %s", s.cfg.App.Name)

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Error : %v", err)

	}
}

func (s *server) httpListening() {
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error : %v", err)
	}
}

func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	s := &server{
		app:        echo.New(),
		db:         db,
		cfg:        cfg,
		middleware: newMiddleware(cfg),
	}

	//Basic Middleware
	//Request Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error : Request Timeout",
		Timeout:      30 * time.Second,
	}))

	//CORS
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//Body Limit
	s.app.Use(middleware.BodyLimit("10M"))

	switch s.cfg.App.Name {
	case "auth":
		s.authService()
	case "player":
		s.playerService()
	case "item":
		s.itemService()
	case "inventory":
		s.inventoryService()
	case "payment":
		s.paymentService()
	}

	//Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s.app.Use(middleware.Logger())

	go s.gracefulShutdown(pctx, quit)

	//Listening
	s.httpListening()
}

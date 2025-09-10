// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"booking-app/config"
	"booking-app/internal/controller/http"
	"booking-app/pkg/httpserver"
	"booking-app/pkg/logger"
	"booking-app/pkg/postgres"
	// TODO: Add these imports when the packages are implemented
	// amqprpc "booking-app/internal/controller/amqp_rpc"
	// "booking-app/internal/repo/persistent"
	// "booking-app/internal/repo/webapi"
	// "booking-app/internal/usecase/translation"
	// "booking-app/pkg/rabbitmq/rmq_rpc/server"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// TODO: Implement use cases and repositories
	// Use-Case
	// translationUseCase := translation.New(
	// 	persistent.New(pg),
	// 	webapi.New(),
	// )

	// TODO: Implement RabbitMQ RPC Server
	// RabbitMQ RPC Server
	// rmqRouter := amqprpc.NewRouter(translationUseCase, l)
	// rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	// if err != nil {
	// 	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	// }

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	// TODO: Pass actual use cases when implemented
	http.NewRouter(httpServer.App, cfg, nil, l)

	// Start servers
	// rmqServer.Start()
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		// TODO: Add RMQ server notification when implemented
		// case err = <-rmqServer.Notify():
		// 	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	// TODO: Add RMQ server shutdown when implemented
	// err = rmqServer.Shutdown()
	// if err != nil {
	// 	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	// }
}

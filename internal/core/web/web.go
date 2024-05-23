package web

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/AnxVit/ozon_1/internal/config"
	"github.com/AnxVit/ozon_1/internal/logger"
	"github.com/AnxVit/ozon_1/internal/repositories"
	"github.com/AnxVit/ozon_1/internal/services"
	"github.com/AnxVit/ozon_1/internal/transport/graphql/api/server"
	"github.com/AnxVit/ozon_1/pkg/postgres"
	"github.com/joho/godotenv"
)

func NewApp() error {
	_ = logger.Init(true)

	if err := godotenv.Load(); err != nil {
		logger.Get().Error("Failed to load .env file: ", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(err.Error()),
		})
		return err
	}

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Get().Error("Failed to load .env file: ", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(err.Error()),
		})
		return err
	}
	postgresClient, err := postgres.NewPostgresClient(cfg)
	if err != nil {
		logger.Get().Error("Failed to load .env file: ", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(err.Error()),
		})
		return err
	}
	habrRepository := repositories.NewHabrRepository(postgresClient)

	habrService := services.NewHabrService(habrRepository)

	errChan := make(chan error)

	svr, err := server.NewAPIServer(cfg, errChan)
	if err != nil {
		logger.Get().Error("Failed to create API server: ", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(err.Error()),
		})
		return err
	}

	server.RegisterAPIHandlers(svr, habrService)

	defer func() {
		logger.Get().Info("Shutting down")
		postgresClient.Close()
		logger.Get().Info("Closed connections")
		close(errChan)
		logger.Get().Info("Closed error channel")
		logger.Get().Info("Exiting")
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	// Ожидание событий в канале ошибок или сигналов
	select {
	case <-ctx.Done():
	case err := <-errChan:
		logger.Get().Error("Error occured: ", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(err.Error()),
		})
	}

	return err
}

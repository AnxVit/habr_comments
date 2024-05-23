package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/AnxVit/ozon_1/api/graphQL/generated"
	"github.com/AnxVit/ozon_1/api/graphQL/interfaces"
	"github.com/AnxVit/ozon_1/internal/config"
	"github.com/AnxVit/ozon_1/internal/logger"
	"github.com/AnxVit/ozon_1/internal/transport/graphql/api"
	"github.com/go-chi/chi/v5"
)

func NewAPIServer(config *config.Config, errChan chan error) (*chi.Mux, error) {
	server := chi.NewRouter()

	address := net.JoinHostPort(config.Server.MainHost, strconv.Itoa(config.Server.MainPort))

	logger.Get().Info("Api server started on")

	go func() {
		if err := http.ListenAndServe(address, server); err != nil {
			errChan <- fmt.Errorf("API server: %w", err)
		}
	}()

	return server, nil
}

func RegisterAPIHandlers(router *chi.Mux, userService api.IHabrService) {
	router.Post("/posts", func(w http.ResponseWriter, r *http.Request) {
		h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
			HabrService: userService,
		}}))
		h.ServeHTTP(w, r)
	})
}

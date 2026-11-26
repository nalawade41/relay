package server

import (
	"log/slog"
	"net/http"

	"github.com/nalawade41/relay/config"
	"github.com/nalawade41/relay/middleware"
	"github.com/nalawade41/relay/router"
	"github.com/nalawade41/relay/store"
)

func NewServer(
	logger *slog.Logger,
	config config.Config,
	eventStore *store.EventStore,
) http.Handler {
	mux := http.NewServeMux()
	router.AddNewRoutes(
		mux,
		logger,
		config,
		eventStore,
	)

	// Middleware that applies to eveery route goes here
	var handler http.Handler = mux
	handler = middleware.RequestLogger(logger, handler)
	return handler
}

package router

import (
	"log/slog"
	"net/http"

	"github.com/nalawade41/relay/config"
	"github.com/nalawade41/relay/handlers"
	"github.com/nalawade41/relay/store"
)

func AddNewRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	config config.Config,
	evenStore *store.EventStore,
) {
	mux.Handle("GET /healthz", handlers.HandleHealthz())
	//Add aditional routes
}

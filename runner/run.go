package runner

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nalawade41/relay/config"
	"github.com/nalawade41/relay/server"
	"github.com/nalawade41/relay/store"
)

func Run(ctx context.Context, args []string, getenv func(string) string, stdin io.Reader, stdout io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// 1. Config: Set up config
	config, err := config.LoadConfig(getenv)
	if err != nil {
		return err
	}

	// 2. Logger: Setup logger
	logger := slog.New(slog.NewJSONHandler(stdout, nil))

	// 3. Database: Create database pool.
	pool, err := pgxpool.New(ctx, config.DatabaseURL)
	if err != nil {
		return fmt.Errorf("create databse pool: %w", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("ping database failed:%w", err)
	}

	// 4. Stores: each one is built from the pool
	eventStore := store.NewEventStore(pool)

	// 5. Server: Crete new server by handing all dependacies
	srv := server.NewServer(
		logger,
		config,
		eventStore,
	)

	// 6. Create the server
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	// 7. Start the server
	go func() {
		logger.Info("listening", "addr", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprint(stdout, "error listening and serving: %s\n", err)
		}
	}()

	// 8. Wait for termination
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(stdout, "error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()
	return err
}

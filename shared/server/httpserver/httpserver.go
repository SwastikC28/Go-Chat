package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

var server *http.Server = nil
var logger = zerolog.Ctx(context.Background()).With().Str("module", "http-server").Logger()

func StartServer(ctx context.Context) {
	config := newHTTPServerConfig()

	server = &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", config.ApiPort),
	}

	logger.Info().Msg(fmt.Sprintf("Starting server at PORT %d", config.ApiPort))
	if err := server.ListenAndServe(); err != nil {
		logger.Err(err).Msg("Error starting server")
		return
	}
}

func StopServer(ctx context.Context) {
	if server != nil {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctxWithTimeout); err != nil {
			logger.Err(err).Msg("Error starting server")
			return
		}

		logger.Info().Msg("Server shutdown successfully.")
	}

}

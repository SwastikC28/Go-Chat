package main

import (
	"context"
	"go-chat/internal/config"
	"os"
	"os/signal"
	_ "shared/logging"
	"syscall"

	"shared/server/httpserver"

	"github.com/rs/zerolog"
)

var service = "go-chat"

func main() {
	ctx := context.Background()
	loggerWithCtx := zerolog.Ctx(ctx).With().Str("service", service).Logger()

	router := config.Initialize()
	go httpserver.StartServer(ctx, router)

	// Listen for termination signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Wait for termination signal
	<-signalCh

	httpserver.StopServer(ctx)
	loggerWithCtx.Info().Msg("Server shutdown successfully")

	os.Exit(0)
}

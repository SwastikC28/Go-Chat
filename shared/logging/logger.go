package logging

import (
	"os"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stdout,
	}).With().Timestamp().Logger()

	zerolog.DefaultContextLogger = &logger

	logger.Debug().Msg("Global Logger Initialized")
}

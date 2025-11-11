package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Log zerolog.Logger

// Init initializes the logger
func Init(env string) {
	zerolog.TimeFieldFormat = time.RFC3339

	if env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	Log = log.Logger
}

// GetLogger returns the global logger instance
func GetLogger() *zerolog.Logger {
	return &Log
}

package service

import (
	"github.com/jaksonkallio/radiate/internal/config"
	"github.com/rs/zerolog"
)

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(config.CurrentConfig.MinLogLevel)
}

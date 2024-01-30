package cmd

import (
	"palworld-chan/pkg/logger"
)

// SetupLogger 开启 Logger
func SetupLogger() error {
	logger.InitLogger()

	return nil
}

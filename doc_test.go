package logger_test

import (
	"fmt"

	"github.com/cloud-spin/logger"
)

func Example() {
	configs := &logger.Configs{
		Enabled: true,
		Level:   logger.LevelInfo,
	}
	logger, err := logger.New(configs)
	if err != nil {
		fmt.Printf("Expected: logger initialized; Got: %s", err.Error())
	}

	logger.Info("Info message with '%s' param", "string")
}

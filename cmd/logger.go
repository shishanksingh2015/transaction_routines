package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

func LoggerConfig() logger.Config {
	fields := map[string]string{
		"time":   "${time}",
		"level":  "INFO",
		"msg":    "request completed",
		"status": "${status}",
		"method": "${method}",
		"path":   "${path}",
	}

	format, err := json.Marshal(fields)
	if err != nil {
		panic(fmt.Errorf("unable to create logger config: %w", err))
	}
	return logger.Config{
		Format:     fmt.Sprintf("%s\n", string(format)),
		TimeFormat: time.RFC3339Nano,
	}
}

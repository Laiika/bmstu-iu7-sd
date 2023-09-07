package main

import (
	"sd/internal/app"
	"sd/internal/config"
	"sd/pkg/logger"
)

func main() {
	log := logger.GetLogger()
	cfg := config.GetConfig(log)

	app.Run(cfg, log)
}

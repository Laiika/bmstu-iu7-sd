package main

import (
	"sd/internal/cli"
	"sd/internal/config"
	"sd/pkg/logger"
)

func main() {
	log := logger.GetLogger()
	cfg := config.GetConfig(log)

	cli.Run(cfg)
}

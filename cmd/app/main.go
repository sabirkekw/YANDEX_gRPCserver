package main

import (
	"fmt"

	"github.com/sabirkekw/YANDEX_gRPCserver/internal/cfg"
	"github.com/sabirkekw/YANDEX_gRPCserver/internal/domain/logger"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()
	logger.Log.Infow("Logger initialized")

	config := cfg.MustLoad()
	logger.Log.Infow("Config loaded", "config", fmt.Sprintf("%+v", config))

	// todo: init database
	// todo: init app
	// todo: start server (app.Run())
	// todo: handle graceful shutdown
}

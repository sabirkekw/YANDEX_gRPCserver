package main

import (
	"fmt"

	"github.com/sabirkekw/YANDEX_gRPCserver/internal/app"
	"github.com/sabirkekw/YANDEX_gRPCserver/internal/cfg"
	"github.com/sabirkekw/YANDEX_gRPCserver/internal/models/order"
	orderservice "github.com/sabirkekw/YANDEX_gRPCserver/internal/services/order"
	"github.com/sabirkekw/YANDEX_gRPCserver/pkg/logger"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()
	logger.Log.Infow("Logger initialized")

	config := cfg.MustLoad()
	logger.Log.Infow("Config loaded\n", "config", fmt.Sprintf("%+v", config))

	db := make(map[string]*order.Order)

	orderService := orderservice.NewService(db, logger.Log)

	application := app.New(logger.Log, config.GRPC.Port, db, orderService)

	if err := application.GRPCServer.Run(); err != nil {
		logger.Log.Fatalw("Failed to run gRPC server", "error", err)
	}
}

package app

import (
	grpcapp "github.com/sabirkekw/YANDEX_gRPCserver/internal/app/grpc"
	grpcserver "github.com/sabirkekw/YANDEX_gRPCserver/internal/grpc/order"
	"github.com/sabirkekw/YANDEX_gRPCserver/internal/models/order"
	"go.uber.org/zap"
)

type App struct {
	GRPCServer *grpcapp.GRPCApp
	Storage    map[string]*order.Order
}

func New(log *zap.SugaredLogger, port int, storage map[string]*order.Order, service grpcserver.OrderService) *App {
	GRPCServer := grpcapp.NewGRPCServer(log, port, service)

	return &App{
		GRPCServer: GRPCServer,
		Storage:    storage,
	}
}

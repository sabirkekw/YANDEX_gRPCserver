package grpcserver

import (
	"context"

	"github.com/sabirkekw/YANDEX_gRPCserver/internal/models/order"
	proto "github.com/sabirkekw/YANDEX_gRPCserver/proto/order"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *order.Order) (string, error)
	GetOrder(ctx context.Context, id string) (*order.Order, error)
	UpdateOrder(ctx context.Context, order *order.Order) error
	DeleteOrder(ctx context.Context, id string) error
	ListOrders(ctx context.Context) ([]*order.Order, error)
}

type Server struct {
	Service OrderService
	Logger *zap.SugaredLogger
	proto.UnimplementedOrderServiceServer
}

func New(service OrderService, logger *zap.SugaredLogger) *Server {
	return &Server{
		Service: service,
		Logger:  logger,
	}
}

func Register(grpc *grpc.Server, server *Server) {
	proto.RegisterOrderServiceServer(grpc, server)
}

func (s *Server) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	const op = "Server.CreateOrder"
	s.Logger.Infow("Received CreateOrder request", "item", req.GetItem(), "quantity", req.GetQuantity(), "op", op)
	order := &order.Order{
		Item:    req.GetItem(),
		Quantity: req.GetQuantity(),
	}
	id, err := s.Service.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &proto.CreateOrderResponse{Id: id}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *proto.GetOrderRequest) (*proto.GetOrderResponse, error) {
	const op = "Server.GetOrder"
	s.Logger.Infow("Received GetOrder request", "id", req.GetId(), "op", op)
	order, err := s.Service.GetOrder(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &proto.GetOrderResponse{
		Order: &proto.Order{
			Id:       order.ID,
			Item:     order.Item,
			Quantity: order.Quantity,
		},
	}, nil
}

func (s *Server) UpdateOrder(ctx context.Context, req *proto.UpdateOrderRequest) (*proto.UpdateOrderResponse, error) {
	const op = "Server.UpdateOrder"
	s.Logger.Infow("Received UpdateOrder request", "id", req.GetId(), "item", req.GetItem(), "quantity", req.GetQuantity(), "op", op)
	order := &order.Order{
		ID:       req.GetId(),
		Item:     req.GetItem(),
		Quantity: req.GetQuantity(),
	}
	if err := s.Service.UpdateOrder(ctx, order); err != nil {
		return nil, err
	}
	return &proto.UpdateOrderResponse{}, nil
}

func (s *Server) DeleteOrder(ctx context.Context, req *proto.DeleteOrderRequest) (*proto.DeleteOrderResponse, error) {
	const op = "Server.DeleteOrder"
	s.Logger.Infow("Received DeleteOrder request", "id", req.GetId(), "op", op)
	if err := s.Service.DeleteOrder(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &proto.DeleteOrderResponse{}, nil
}

func (s *Server) ListOrders(ctx context.Context, req *proto.ListOrdersRequest) (*proto.ListOrdersResponse, error) {
	const op = "Server.ListOrders"
	s.Logger.Infow("Received ListOrders request", "op", op)
	orders, err := s.Service.ListOrders(ctx)
	if err != nil {
		return nil, err
	}
	var protoOrders []*proto.Order
	for _, order := range orders {
		protoOrders = append(protoOrders, &proto.Order{
			Id:       order.ID,
			Item:     order.Item,
			Quantity: order.Quantity,
		})
	}
	return &proto.ListOrdersResponse{Orders: protoOrders}, nil
}

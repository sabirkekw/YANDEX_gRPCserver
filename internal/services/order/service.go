package orderservice

import (
	"context"
	"sync"

	order "github.com/sabirkekw/YANDEX_gRPCserver/internal/models/order"
	"go.uber.org/zap"
	uuid "github.com/google/uuid"
)

type Service struct {
	storage map[string]string
	logger  *zap.SugaredLogger
}

func NewService(storage map[string]*order.Order, logger *zap.SugaredLogger) *Service {
	return &Service{storage: storage, logger: logger}
}

func (s *Service) CreateOrder(ctx context.Context, order *order.Order) (string, error) {
	const op = "Service.CreateOrder"
	s.logger.Infow("Creating order", "item", order.Item, "quantity", order.Quantity, "op", op)

	id := uuid.New().String()
	order.ID = id

	mx := sync.Mutex{}
	mx.Lock()
	s.storage[id] = order.Item
	mx.Unlock()
	return id, nil
}

func (s *Service) GetOrder(ctx context.Context, id string) (*order.Order, error) {
	const op = "Service.GetOrder"
	s.logger.Infow("Getting order", "id", id, "op", op)
	
	order, exists := s.storage[id]
	if !exists {
		return nil, nil
	}
	return &order.Order{ID: id, Item: order.}, nil
}

func (s *Service) UpdateOrder(ctx context.Context, order *order.Order) error {
	const op = "Service.UpdateOrder"
	s.logger.Infow("Updating order", "id", order.ID, "item", order.Item, "quantity", order.Quantity, "op", op)
	panic("implement me")
}

func (s *Service) DeleteOrder(ctx context.Context, id string) error {
	const op = "Service.DeleteOrder"
	s.logger.Infow("Deleting order", "id", id, "op", op)
	panic("implement me")
}

func (s *Service) ListOrders(ctx context.Context) ([]*order.Order, error) {
	const op = "Service.ListOrders"
	s.logger.Infow("Listing orders", "op", op)
	panic("implement me")
}

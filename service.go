package main

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"

	ordersvc "github.com/shijuvar/gokit-examples/services/order"
)

// service implements the Order Service
type service struct {
	repository ordersvc.Repository
	logger     log.Logger
}
type Service interface {
	Create(ctx context.Context, order Order) (string, error)
}

var (
	ErrOrderNotFound   = errors.New("order not found")
	ErrCmdRepository   = errors.New("unable to command repository")
	ErrQueryRepository = errors.New("unable to query repository")
)

type orderService struct{}

// NewService creates and returns a new Order service instance
func NewService(rep ordersvc.Repository, logger log.Logger) ordersvc.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// Create makes an order
func (s *service) Create(ctx context.Context, order ordersvc.Order) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	order.ID = id
	order.Status = "Pending"
	order.CreatedOn = time.Now().Unix()

	if err := s.repository.CreateOrder(ctx, order); err != nil {
		level.Error(logger).Log("err", err)
		return "", ordersvc.ErrCmdRepository
	}
	return id, nil
}

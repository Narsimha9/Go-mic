package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type OrderService interface {
	Create(ctx context.Context, order Order) (string, error)
}

var (
	ErrOrderNotFound   = errors.New("order not found")
	ErrCmdRepository   = errors.New("unable to command repository")
	ErrQueryRepository = errors.New("unable to query repository")
)

type orderService struct{}

func (orderService) Create(ctx context.Context, order Order) (string, error) {
	fmt.Println("Crete method implementation called")
	// logger := log.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	order.ID = id
	order.Status = "Pending"
	order.CreatedOn = time.Now().Unix()

	// if err := s.repository.CreateOrder(ctx, order); err != nil {
	// 	level.Error(logger).Log("err", err)
	// 	return "", ordersvc.ErrCmdRepository
	// }
	return id, nil
	// return order, nil
}

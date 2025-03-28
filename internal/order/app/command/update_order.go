package command

import (
	"context"
	"github.com/Oscar-Lu-01/gorder/common/decorator"
	domain "github.com/Oscar-Lu-01/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

type UpdateOrder struct {
	order    *domain.Order
	UpdateFn func(context.Context, *domain.Order) (*domain.Order, error)
}

type UpdateOrderHandler decorator.CommandHandler[UpdateOrder, interface{}]

func NewUpdateOrderHandler(
	orderRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) UpdateOrderHandler {
	if orderRepo == nil {
		panic("nil orderRepo")
	}
	return decorator.ApplyCommandDecorators[UpdateOrder, interface{}](
		updateOrderHandler{orderRepo}, logger, metricsClient)
}

type updateOrderHandler struct {
	orderRepo domain.Repository
}

func (c updateOrderHandler) Handle(ctx context.Context, cmd UpdateOrder) (interface{}, error) {
	//校验
	if cmd.UpdateFn == nil {
		logrus.Warning("UpdateOrderHandler got nil update function, order = %#v", cmd.order.ID)
		//如果函数是nil就什么都不管
		cmd.UpdateFn = func(ctx context.Context, order *domain.Order) (*domain.Order, error) { return order, nil }
	}
	err := c.orderRepo.Update(ctx, cmd.order, cmd.UpdateFn)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

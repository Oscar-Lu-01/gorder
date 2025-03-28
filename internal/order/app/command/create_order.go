package command

import (
	"context"
	"github.com/Oscar-Lu-01/gorder/common/decorator"
	"github.com/Oscar-Lu-01/gorder/common/genproto/orderpb"
	domain "github.com/Oscar-Lu-01/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

type CreateOrder struct {
	CustomerID string
	Items      []*orderpb.ItemWithQuantity
}

type CreateOrderResult struct {
	OrderID string
}

type CreateOrderHandler decorator.CommandHandler[CreateOrder, *CreateOrderResult]

func NewCreateOrderHandler(
	orderRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CreateOrderHandler {
	if orderRepo == nil {
		panic("nil orderRepo")
	}
	return decorator.ApplyCommandDecorators[CreateOrder, *CreateOrderResult](
		createOrderHandler{orderRepo}, logger, metricsClient)
}

type createOrderHandler struct {
	orderRepo domain.Repository
	//stockGRPC
}

func (h createOrderHandler) Handle(ctx context.Context, cmd CreateOrder) (*CreateOrderResult, error) {
	//TODO:call stock grpc to get items
	//模拟
	var stockResponse []*orderpb.Item
	for _, item := range cmd.Items {
		stockResponse = append(stockResponse, &orderpb.Item{
			Quantity: item.Quantity,
			ID:       item.ID,
		})
	}
	order, err := h.orderRepo.Create(ctx, &domain.Order{
		CustomerID: cmd.CustomerID,
		Items:      stockResponse,
	})
	if err != nil {
		return nil, err
	}
	return &CreateOrderResult{OrderID: order.ID}, nil
}

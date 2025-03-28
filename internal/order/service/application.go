package service

import (
	"context"
	"github.com/Oscar-Lu-01/gorder/common/metrics"
	"github.com/Oscar-Lu-01/gorder/order/adapters"
	"github.com/Oscar-Lu-01/gorder/order/app"
	"github.com/Oscar-Lu-01/gorder/order/app/command"
	"github.com/Oscar-Lu-01/gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderInmemRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.ToDoMetrics{}
	return app.Application{
		Commands: app.Commands{
			CreateOrder: command.NewCreateOrderHandler(orderInmemRepo, logger, metricsClient),
			UpdateOrder: command.NewUpdateOrderHandler(orderInmemRepo, logger, metricsClient),
		},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderInmemRepo, logger, metricsClient),
		},
	}
}

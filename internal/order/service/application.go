package service

import (
	"context"
	"github.com/Oscar-Lu-01/gorder/common/metrics"
	"github.com/Oscar-Lu-01/gorder/order/adapters"
	"github.com/Oscar-Lu-01/gorder/order/app"
	"github.com/Oscar-Lu-01/gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderInmemRepo := adapters.NewMemoryOrderRepository()
	logrus := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.ToDoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderInmemRepo, logrus, metricsClient),
		},
	}
}

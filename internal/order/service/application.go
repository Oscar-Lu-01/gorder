package service

import (
	"context"
	"github.com/Oscar-Lu-01/gorder/order/adapters"
	"github.com/Oscar-Lu-01/gorder/order/app"
)

func NewApplication(ctx context.Context) app.Application {

	orderRepo := adapters.NewMemoryOrderRepository()
	return app.Application{}
}

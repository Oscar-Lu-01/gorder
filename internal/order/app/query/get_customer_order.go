package query

import (
	"context"
	"github.com/Oscar-Lu-01/gorder/common/decorator"
	domain "github.com/Oscar-Lu-01/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

// 参照QueryHandler,需要一个query和一个result
// order结构体
type GetCustomerOrder struct {
	CustomerID string
	OrderID    string
}

// 泛型接口QueryHandler的自定义接口
type GetCustomerOrderHandler decorator.QueryHandler[GetCustomerOrder, *domain.Order]

func NewGetCustomerOrderHandler(
	orderRepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) GetCustomerOrderHandler {
	//没检测到仓库，不需要初始化
	if orderRepo == nil {
		panic("nil orderRepo")
	}
	return decorator.ApplyQueryDecorators[GetCustomerOrder, *domain.Order](
		getCustomerOrderHandler{orderRepo: orderRepo},
		logger,
		metricClient)

}

// 数据库接口在domain，实现在adapter
// 依赖domain的repository，更换数据库只需要更改adapter中的实现
type getCustomerOrderHandler struct {
	orderRepo domain.Repository
}

func (g getCustomerOrderHandler) Handle(ctx context.Context, query GetCustomerOrder) (*domain.Order, error) {
	order, err := g.orderRepo.Get(ctx, query.OrderID, query.CustomerID)
	if err != nil {
		return nil, err
	} else {
		return order, nil
	}
}

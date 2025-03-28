package decorator

import (
	"context"
	"github.com/sirupsen/logrus"
)

// QueryHandler defines a generic type that receives a commend C,returns a result R
// 泛型接口-》自定义接口-》自定义实现
type QueryHandler[C, R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}

// 增强顺序：日志最先
func ApplyQueryDecorators[H, R any](handler QueryHandler[H, R], logger *logrus.Entry, metricsClient MetricsClient) QueryHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		logger: logger,
		base: queryMetricsDecorator[H, R]{
			client: metricsClient,
			base:   handler,
		},
	}
}

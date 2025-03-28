package decorator

import (
	"context"
	"github.com/sirupsen/logrus"
)

// QueryHandler defines a generic type that receives a commend C,returns a result R
// 泛型接口-》自定义接口-》自定义实现
type CommandHandler[C, R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}

// 增强顺序：日志最先
func ApplyCommandDecorators[C, R any](handler CommandHandler[C, R], logger *logrus.Entry, metricsClient MetricsClient) CommandHandler[C, R] {
	return queryLoggingDecorator[C, R]{
		logger: logger,
		base: queryMetricsDecorator[C, R]{
			client: metricsClient,
			base:   handler,
		},
	}
}

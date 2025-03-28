package decorator

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

// 泛型结构体
type queryLoggingDecorator[C, R any] struct {
	logger *logrus.Entry
	base   QueryHandler[C, R]
}

func (q queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	logger := q.logger.WithFields(logrus.Fields{
		"query":      cmd,
		"query_body": fmt.Sprintf("%#v", cmd),
	})
	logger.Debug("Executing query")
	//小技巧，使用defer打印结果
	defer func() {
		if err == nil {
			logger.WithError(err).Warn("Query executed successfully")
		} else {
			logger.WithError(err).Warn("Failed to executed query")
		}
	}()
	return q.base.Handle(ctx, cmd)
}

// 提取query.XXXXHandler
func generateActionName(cmd any) string {
	return strings.Split(fmt.Sprintf("%T", cmd), ".")[1]
}

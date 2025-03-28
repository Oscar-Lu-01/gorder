package decorator

import (
	"context"
	"fmt"
	"strings"
	"time"
)

//额外实现一些指标的装饰器，如：计时、函数调用次数

// 需要依赖第三方服务上报指标，目前暂时写入内存
type MetricsClient interface {
	Inc(key string, value int)
}

type queryMetricsDecorator[C, R any] struct {
	client MetricsClient
	base   QueryHandler[C, R]
}

func (q queryMetricsDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	start := time.Now()
	actionName := strings.ToLower(generateActionName(cmd))
	//小技巧，使用defer打印结果
	defer func() {
		end := time.Since(start)
		q.client.Inc(fmt.Sprintf("querys.%s.duration", actionName), int(end.Seconds()))
		if err == nil {
			q.client.Inc(fmt.Sprintf("querys.%s.success", actionName), 1)
		} else {
			q.client.Inc(fmt.Sprintf("querys.%s.failure", actionName), 1)
		}
	}()
	return q.base.Handle(ctx, cmd)
}

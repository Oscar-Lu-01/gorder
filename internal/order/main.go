package main

import (
	"context"
	"github.com/Oscar-Lu-01/gorder/common/config"
	"github.com/Oscar-Lu-01/gorder/common/genproto/orderpb"
	"github.com/Oscar-Lu-01/gorder/common/server"
	"github.com/Oscar-Lu-01/gorder/order/ports"
	"github.com/Oscar-Lu-01/gorder/order/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// 编译时候最先执行
func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")

	//在service/application.go中实现 application
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)

	//在携程里运行GRPC服务，不会阻塞http
	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer(application)
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	server.RunHttpServer(serviceName, func(router *gin.Engine) {
		//写的和第六集10min左右有一点不一样
		svc := NewHttpServer(application)
		ports.RegisterHandlersWithOptions(router, svc, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})

}

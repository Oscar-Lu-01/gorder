package main

import (
	"github.com/Oscar-Lu-01/gorder/common/config"
	"github.com/Oscar-Lu-01/gorder/common/genproto/orderpb"
	"github.com/Oscar-Lu-01/gorder/common/server"
	"github.com/Oscar-Lu-01/gorder/order/ports"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

// 编译时候最先执行
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")

	//在携程里运行GRPC服务，不会阻塞http
	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer()
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	server.RunHttpServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HttpServer{}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})

}

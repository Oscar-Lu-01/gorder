package main

import (
	"context"
	"github.com/Oscar-Lu-01/gorder/common/config"
	"github.com/Oscar-Lu-01/gorder/common/genproto/stockpb"
	"github.com/Oscar-Lu-01/gorder/common/server"
	"github.com/Oscar-Lu-01/gorder/stock/ports"
	"github.com/Oscar-Lu-01/gorder/stock/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// 编译时候最先执行,初始化viper，不然会拿不到server-to-run
func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	ServerType := viper.GetString("stock.server-to-run")

	//在service/application.go中实现 application
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)

	switch ServerType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			stockpb.RegisterStockServiceServer(server, ports.NewGRPCServer(application))
		})
	case "http":
		//暂时不用
	default:
		panic("unexpected server type")
	}

}

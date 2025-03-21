package main

import (
	"github.com/Oscar-Lu-01/gorder/common/genproto/stockpb"
	"github.com/Oscar-Lu-01/gorder/common/server"
	"github.com/Oscar-Lu-01/gorder/stock/ports"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	serviceName := viper.GetString("stock.service-name")
	ServerType := viper.GetString("stock.server-to-run")
	switch ServerType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			stockpb.RegisterStockServiceServer(server, ports.NewGRPCServer())
		})
	case "http":
		//暂时不用
	default:
		panic("unexpected server type")
	}

}

package server

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

// 初始化logrusEntry
func init() {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)
	grpc_logrus.ReplaceGrpcLogger(logrus.NewEntry(logger))
}

func RunGRPCServer(serverName string, registerServer func(server *grpc.Server)) {
	addr := viper.Sub(serverName).GetString("grpc-addr")
	if addr == "" {
		addr = viper.GetString("fallback-grpc-addr")

	}
	RunGRPCServerOnAddr(addr, registerServer)
}

func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	//参照https://github.com/grpc-ecosystem/go-grpc-middleware
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_tags.UnaryServerInterceptor(grpc_tags.WithFieldExtractor(grpc_tags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			//srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			//logging.UnaryServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
			//selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authFn), selector.MatchFunc(allButHealthZ)),
			//recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			grpc_tags.StreamServerInterceptor(grpc_tags.WithFieldExtractor(grpc_tags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry),
			//srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			//logging.StreamServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
			//selector.StreamServerInterceptor(auth.StreamServerInterceptor(authFn), selector.MatchFunc(allButHealthZ)),
			//recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	)
	registerServer(grpcServer)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Panic(err)
	}
	//logrus.WithField("addr", addr).Info("Listening on " + addr)
	logrus.Infof("grpc server listening on %s", addr)
	if err := grpcServer.Serve(listener); err != nil {
		logrus.Panic(err)
	}
}

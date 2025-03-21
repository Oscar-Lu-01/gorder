package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// func(router)为匿名函数
func RunHttpServer(serviceName string, wrapper func(router *gin.Engine)) {
	addr := viper.Sub(serviceName).GetString("http-addr")
	RunHttpServerOnAddr(addr, wrapper)
}

func RunHttpServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRouter := gin.New()
	wrapper(apiRouter)
	apiRouter.Group("/api")
	//logrus.WithField("addr", addr).Info("start http server")
	logrus.Infof("http server listening on %s", addr)
	if err := apiRouter.Run(addr); err != nil {
		panic(err)
	}
}

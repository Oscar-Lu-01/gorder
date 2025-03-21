package main

import (
	"github.com/Oscar-Lu-01/gorder/order/app"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) *HttpServer {
	return &HttpServer{app: app}
}

func (h HttpServer) PostCustomerCustomerIDOrders(c *gin.Context, customerID string) {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) GetCustomerCustomerIDOrdersOrderID(c *gin.Context, customerID string, orderID string) {
	//TODO implement me
	panic("implement me")
}

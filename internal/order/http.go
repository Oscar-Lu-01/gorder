package main

import (
	"github.com/Oscar-Lu-01/gorder/order/app"
	"github.com/Oscar-Lu-01/gorder/order/app/query"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) *HttpServer {
	return &HttpServer{app: app}
}

func (h HttpServer) PostCustomerCustomerIDOrders(ctx *gin.Context, customerID string) {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) GetCustomerCustomerIDOrdersOrderID(ctx *gin.Context, customerID string, orderID string) {
	order, err := h.app.Queries.GetCustomerOrder.Handle(ctx, query.GetCustomerOrder{
		CustomerID: "fake-ID",
		OrderID:    "fake-customer-id",
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": order})
}

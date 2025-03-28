package main

import (
	"github.com/Oscar-Lu-01/gorder/common/genproto/orderpb"
	"github.com/Oscar-Lu-01/gorder/order/app"
	"github.com/Oscar-Lu-01/gorder/order/app/command"
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
	var require orderpb.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&require); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	res, err := h.app.Commands.CreateOrder.Handle(ctx, command.CreateOrder{
		CustomerID: customerID,
		Items:      require.Items,
	})
	if err != nil {
		//反正视频这么写的
		ctx.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"customer_id": require.CustomerID,
		"order_id":    res.OrderID})
}

func (h HttpServer) GetCustomerCustomerIDOrdersOrderID(ctx *gin.Context, customerID string, orderID string) {
	order, err := h.app.Queries.GetCustomerOrder.Handle(ctx, query.GetCustomerOrder{
		CustomerID: customerID,
		OrderID:    orderID,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": order})
}

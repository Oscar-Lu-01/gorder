package order

import "github.com/Oscar-Lu-01/gorder/common/genproto/orderpb"

type Order struct {
	ID          string
	CustomerID  string
	Status      string
	PaymentLink string
	Items       []*orderpb.Item
}

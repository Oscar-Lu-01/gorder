package adapters

import (
	"context"
	domain "github.com/Oscar-Lu-01/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

// 实现
type MemoryOrderRepository struct {
	//共享锁，所以传指针
	lock  *sync.RWMutex
	store []*domain.Order
}

// 测试代码：data类型存储在domain中，从domain里取
var fakeData = []*domain.Order{}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	//测试代码
	s := make([]*domain.Order, 0)
	s = append(s, &domain.Order{
		ID:          "fake-ID",
		CustomerID:  "fake-customer-id",
		Status:      "fake-status",
		PaymentLink: "fake-payment-link",
		Items:       nil,
	})
	return &MemoryOrderRepository{
		lock:  &sync.RWMutex{},
		store: make([]*domain.Order, 0),
	}
}

func (m MemoryOrderRepository) Create(_ context.Context, order *domain.Order) (*domain.Order, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	NewOrder := &domain.Order{
		ID:          strconv.FormatInt(time.Now().Unix(), 10),
		CustomerID:  order.CustomerID,
		Status:      order.Status,
		PaymentLink: order.PaymentLink,
		Items:       order.Items,
	}
	m.store = append(m.store, order)
	logrus.WithFields(logrus.Fields{
		"input_order":        order,
		"store_after_create": m.store,
	}).Debugf("memory_order_repo_create")
	return NewOrder, nil
}

func (m MemoryOrderRepository) Get(_ context.Context, id, customerID string) (*domain.Order, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _, order := range m.store {
		if order.ID == id && order.CustomerID == customerID {
			return order, nil
		}
	}
	return nil, domain.NotFoundError{OrderID: id}
}

func (m MemoryOrderRepository) Update(ctx context.Context, o *domain.Order, updateFn func(context.Context, *domain.Order) (*domain.Order, error)) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	found := false
	for i, order := range m.store {
		if order.ID == o.ID && order.CustomerID == o.CustomerID {
			found = true
			updateOrder, err := updateFn(ctx, o)
			if err != nil {
				return err
			}
			m.store[i] = updateOrder
		}
	}
	if !found {
		return domain.NotFoundError{OrderID: o.ID}
	}
	return nil
}

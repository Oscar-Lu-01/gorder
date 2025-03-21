package adapters

import (
	"context"
	"github.com/Oscar-Lu-01/gorder/common/genproto/orderpb"
	domain "github.com/Oscar-Lu-01/gorder/stock/domain/stock"
	"sync"
)

type MemoryStockRepository struct {
	//共享锁，所以传指针
	lock  *sync.RWMutex
	store map[string]*orderpb.Item
}

// 暂时写死
var stub = map[string]*orderpb.Item{
	"item_id": {
		ID:       "foo_item",
		Name:     "stub_item",
		Quantity: 10000,
		PriceID:  "stub_item_price_id",
	},
}

func NewMemoryStockRepository() *MemoryStockRepository {
	return &MemoryStockRepository{
		lock:  &sync.RWMutex{},
		store: stub,
	}
}

func (m MemoryStockRepository) GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var (
		res     []*orderpb.Item
		missing []string
	)
	for _, id := range ids {
		if item, exist := m.store[id]; exist {
			res = append(res, item)
		} else {
			missing = append(missing, id)
		}
	}
	if len(res) == len(ids) {
		return res, nil
	}
	//思考传值为什么用Missing_id: missing
	return nil, domain.NotFoundError{Missing_id: missing}
}

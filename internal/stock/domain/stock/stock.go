package stock

import (
	"context"
	"fmt"
	"github.com/Oscar-Lu-01/gorder/common/genproto/orderpb"
	"strings"
)

type Repository interface {
	GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error)
}

type NotFoundError struct {
	Missing_id []string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("These items not found: %s", strings.Join(e.Missing_id, ","))
}

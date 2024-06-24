package mem

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/internal/service"
)

func (r *MemDb) WithTransaction(ctx context.Context, fn service.TxFunc) error {
	err := fn(ctx)
	return err
}

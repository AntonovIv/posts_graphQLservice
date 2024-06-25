package postgre

import (
	"context"

	postgreMr "github.com/AntonovIv/post_graphQlservice/internal/db/postgre"
	"github.com/AntonovIv/post_graphQlservice/internal/service"
)

type repository struct {
	db postgreMr.QueryManager
}

func New(db postgreMr.QueryManager) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) WithTransaction(ctx context.Context, fn service.TxFunc) error {
	return r.db.WithTransaction(ctx, fn)
}

//go:generate mockgen -source=internal\service\posts\service.go -destination=internal\repository\postgre\mock\postgre_mock.go

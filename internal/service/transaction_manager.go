package service

import "context"

type TxFunc func(ctx context.Context) error

type TransactionManager interface {
	WithTransaction(ctx context.Context, fn TxFunc) error
}

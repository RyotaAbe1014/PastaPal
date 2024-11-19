package transaction

import "context"

type TransactionManager interface {
	RunInTransaction(ctx context.Context, f func(ctx context.Context) error) error
}

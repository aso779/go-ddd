package infrastructure

type TransactionManager[T any] interface {
	Commit()
	GetTransaction() T
	Rollback()
}

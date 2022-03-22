package infrastructure

import (
	"context"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type CrudRepository[E metadata.Entity, T any] interface {
	FindOne(
		ctx context.Context,
		tx TransactionManager[T],
		fields []string,
		spec dataset.Specifier,
	) (*E, error)

	FindOneById(
		ctx context.Context,
		tx TransactionManager[T],
		fields []string,
		id metadata.PrimaryKey,
	) (*E, error)

	FindAll(
		ctx context.Context,
		tx TransactionManager[T],
		fields []string,
		spec dataset.Specifier,
		page dataset.Pager,
		sort dataset.Sorter,
	) (*[]E, error)

	FindAllByIds(
		ctx context.Context,
		tx TransactionManager[T],
		fields []string,
		ids []metadata.PrimaryKey,
	) (*[]E, error)

	Count(
		ctx context.Context,
		tx TransactionManager[T],
		spec dataset.Specifier,
	) (int, error)

	Create(
		ctx context.Context,
		tx TransactionManager[T],
		ent *E,
		fields []string,
	) (*E, error)

	Update(
		ctx context.Context,
		tx TransactionManager[T],
		ent *E,
		fields []string,
		ftu []string,
	) (*E, error)

	Delete(
		ctx context.Context,
		tx TransactionManager[T],
		spec dataset.Specifier,
	) (int, error)
}

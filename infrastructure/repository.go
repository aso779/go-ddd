package infrastructure

import (
	"context"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type CrudRepository[T metadata.Entity] interface {
	FindOne(
		ctx context.Context,
		tx TransactionManager,
		fields []string,
		spec dataset.Specifier,
	) (*T, error)

	FindOneById(
		ctx context.Context,
		tx TransactionManager,
		fields []string,
		id metadata.PrimaryKey,
	) (*T, error)

	FindAll(
		ctx context.Context,
		tx TransactionManager,
		fields []string,
		spec dataset.Specifier,
		page dataset.Pager,
		sort dataset.Sorter,
	) (*[]T, error)

	FindAllByIds(
		ctx context.Context,
		tx TransactionManager,
		fields []string,
		ids []metadata.PrimaryKey,
	) (*[]T, error)

	Count(
		ctx context.Context,
		tx TransactionManager,
		spec dataset.Specifier,
	) (int, error)

	Create(
		ctx context.Context,
		tx TransactionManager,
		ent *T,
		fields []string,
	) (*T, error)

	Update(
		ctx context.Context,
		tx TransactionManager,
		ent *T,
		fields []string,
		ftu []string,
	) (*T, error)

	Delete(
		ctx context.Context,
		tx TransactionManager,
		spec dataset.Specifier,
	) (int, error)
}

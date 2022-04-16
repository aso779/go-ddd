package infrastructure

import (
	"context"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type CrudRepository[E metadata.Entity, T any] interface {
	FindOne(
		ctx context.Context,
		tx any,
		fields []string,
		spec dataset.Specifier,
	) (*E, error)

	FindOneByPk(
		ctx context.Context,
		tx any,
		fields []string,
		pk metadata.PrimaryKey,
	) (*E, error)

	FindAll(
		ctx context.Context,
		tx any,
		fields []string,
		spec dataset.Specifier,
	) (*[]E, error)

	FindPage(
		ctx context.Context,
		tx any,
		fields []string,
		spec dataset.Specifier,
		page dataset.Pager,
		sort dataset.Sorter,
	) (*[]E, error)

	FindAllByPks(
		ctx context.Context,
		tx any,
		fields []string,
		pks []metadata.PrimaryKey,
	) (*[]E, error)

	Count(
		ctx context.Context,
		tx any,
		spec dataset.Specifier,
	) (int, error)

	CreateOne(
		ctx context.Context,
		tx any,
		ent *E,
		fields []string,
	) (*E, error)

	UpdateOne(
		ctx context.Context,
		tx any,
		ent *E,
		fields []string,
	) (*E, error)

	Delete(
		ctx context.Context,
		tx any,
		spec dataset.Specifier,
	) (int, error)
}

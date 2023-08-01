package dataspec

import (
	"fmt"

	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type NotIsNullSpecification struct {
	field Field
}

func NewNotIsNull(field string) dataset.Specifier {
	return &NotIsNullSpecification{
		field: NewField(field),
	}
}

func (r *NotIsNullSpecification) Joins(meta metadata.Meta) []metadata.Join {
	return join(meta, r.field)
}

func (r *NotIsNullSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s IS NOT NULL", r.field.ColumnName(meta))
}

func (r *NotIsNullSpecification) Values() []any {
	return []any{}
}

func (r *NotIsNullSpecification) IsEmpty() bool {
	return false
}

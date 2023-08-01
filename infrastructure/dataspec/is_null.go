package dataspec

import (
	"fmt"

	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type IsNullSpecification struct {
	field Field
}

func NewIsNull(field string) dataset.Specifier {
	return &IsNullSpecification{
		field: NewField(field),
	}
}

func (r *IsNullSpecification) Joins(meta metadata.Meta) []metadata.Join {
	return join(meta, r.field)
}

func (r *IsNullSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s IS NULL", r.field.ColumnName(meta))
}

func (r *IsNullSpecification) Values() []any {
	return []any{}
}

func (r *IsNullSpecification) IsEmpty() bool {
	return false
}

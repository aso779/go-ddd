package dataspec

import (
	"fmt"

	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type NotEqualSpecification struct {
	field Field
	value any
}

func NewNotEqual(field string, value any) dataset.Specifier {
	return &NotEqualSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *NotEqualSpecification) Joins(meta metadata.Meta) []metadata.Join {
	return join(meta, r.field)
}

func (r *NotEqualSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s != ?", r.field.ColumnName(meta))
}

func (r *NotEqualSpecification) Values() []any {
	return []any{r.value}
}

func (r *NotEqualSpecification) IsEmpty() bool {
	return r.value == nil
}

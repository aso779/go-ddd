package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type NotEqualSpecification struct {
	field Field
	value interface{}
}

func NewNotEqual(field string, value interface{}) dataset.Specifier {
	return &NotEqualSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *NotEqualSpecification) Joins(meta metadata.Meta) []string {
	return join(meta, r.field)
}

func (r NotEqualSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s != ?", r.field.ColumnName(meta))
}

func (r NotEqualSpecification) Values() []interface{} {
	return []interface{}{r.value}
}

func (r *NotEqualSpecification) IsEmpty() bool {
	return r.value == nil
}

package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type EqualSpecification struct {
	field Field
	value interface{}
}

func NewEqual(field string, value interface{}) dataset.Specifier {
	return &EqualSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *EqualSpecification) Joins(meta metadata.Meta) []string {
	return join(meta, r.field)
}

func (r *EqualSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s = ?", r.field.ColumnName(meta))
}

func (r *EqualSpecification) Values() []interface{} {
	return []interface{}{r.value}
}

func (r *EqualSpecification) IsEmpty() bool {
	return r.value == nil
}

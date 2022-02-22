package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type InSpecification struct {
	field Field
	value interface{}
}

func NewIn(field string, value interface{}) dataset.Specifier {
	return &InSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *InSpecification) Joins(meta metadata.Meta) []string {
	return join(meta, r.field)
}

func (r InSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s IN (?)", r.field.ColumnName(meta))
}

func (r InSpecification) Values() []interface{} {
	return []interface{}{r.value}
}

func (r *InSpecification) IsEmpty() bool {
	return r.value == nil
}

package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type GteSpecification struct {
	field Field
	value interface{}
}

func NewGte(field string, value interface{}) dataset.Specifier {
	return &GteSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *GteSpecification) Joins(meta metadata.Meta) []string {
	return join(meta, r.field)
}

func (r *GteSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s >= ?", r.field.ColumnName(meta))
}

func (r *GteSpecification) Values() []interface{} {
	return []interface{}{r.value}
}

func (r *GteSpecification) IsEmpty() bool {
	return r.value == nil
}

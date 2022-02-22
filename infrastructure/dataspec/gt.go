package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type GtSpecification struct {
	field Field
	value interface{}
}

func NewGt(field string, value interface{}) dataset.Specifier {
	return &GtSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *GtSpecification) Joins(meta metadata.Meta) []string {
	return join(meta, r.field)
}

func (r *GtSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s > ?", r.field.ColumnName(meta))
}

func (r *GtSpecification) Values() []interface{} {
	return []interface{}{r.value}
}

func (r *GtSpecification) IsEmpty() bool {
	return r.value == nil
}

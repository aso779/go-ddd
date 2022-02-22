package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type LteSpecification struct {
	field Field
	value interface{}
}

func NewLte(field string, value interface{}) dataset.Specifier {
	return &LteSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *LteSpecification) Joins(meta metadata.Meta) []string {
	return join(meta, r.field)
}

func (r *LteSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s <= ?", r.field.ColumnName(meta))
}

func (r *LteSpecification) Values() []interface{} {
	return []interface{}{r.value}
}

func (r *LteSpecification) IsEmpty() bool {
	return r.value == nil
}

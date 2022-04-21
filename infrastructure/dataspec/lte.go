package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type LteSpecification struct {
	field Field
	value any
}

func NewLte(field string, value any) dataset.Specifier {
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

func (r *LteSpecification) Values() []any {
	return []any{r.value}
}

func (r *LteSpecification) IsEmpty() bool {
	return r.value == nil
}

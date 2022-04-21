package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type LtSpecification struct {
	field Field
	value any
}

func NewLt(field string, value any) dataset.Specifier {
	return &LtSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *LtSpecification) Joins(meta metadata.Meta) []string {
	return join(meta, r.field)
}

func (r *LtSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s < ?", r.field.ColumnName(meta))
}

func (r *LtSpecification) Values() []any {
	return []any{r.value}
}

func (r *LtSpecification) IsEmpty() bool {
	return r.value == nil
}

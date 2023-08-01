package dataspec

import (
	"fmt"

	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type GtSpecification struct {
	field Field
	value any
}

func NewGt(field string, value any) dataset.Specifier {
	return &GtSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *GtSpecification) Joins(meta metadata.Meta) []metadata.Join {
	return join(meta, r.field)
}

func (r *GtSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s > ?", r.field.ColumnName(meta))
}

func (r *GtSpecification) Values() []any {
	return []any{r.value}
}

func (r *GtSpecification) IsEmpty() bool {
	return r.value == nil
}

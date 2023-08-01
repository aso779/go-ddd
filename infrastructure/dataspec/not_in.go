package dataspec

import (
	"fmt"

	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type NotInSpecification struct {
	field Field
	value any
}

func NewNotIn(field string, value any) dataset.Specifier {
	return &InSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *NotInSpecification) Joins(meta metadata.Meta) []metadata.Join {
	return join(meta, r.field)
}

func (r *NotInSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s NOT IN (?)", r.field.ColumnName(meta))
}

func (r *NotInSpecification) Values() []any {
	return []any{r.value}
}

func (r *NotInSpecification) IsEmpty() bool {
	return r.value == nil
}

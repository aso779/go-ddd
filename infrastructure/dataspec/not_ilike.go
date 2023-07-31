package dataspec

import (
	"fmt"

	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type NotILikeSpecification struct {
	field Field
	value string
}

func NewNotILike(field string, value string) dataset.Specifier {
	return &NotILikeSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *NotILikeSpecification) Joins(meta metadata.Meta) []string {
	return join(meta, r.field)
}

func (r *NotILikeSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s NOT ILIKE ?", r.field.ColumnName(meta))
}

func (r *NotILikeSpecification) Values() []any {
	return []any{r.value}
}

func (r *NotILikeSpecification) IsEmpty() bool {
	return r.value == ""
}

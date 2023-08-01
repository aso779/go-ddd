package dataspec

import (
	"fmt"

	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type NotLikeSpecification struct {
	field Field
	value string
}

func NewNotLike(field string, value string) dataset.Specifier {
	return &NotLikeSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *NotLikeSpecification) Joins(meta metadata.Meta) []metadata.Join {
	return join(meta, r.field)
}

func (r *NotLikeSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s NOT LIKE ?", r.field.ColumnName(meta))
}

func (r *NotLikeSpecification) Values() []any {
	return []any{r.value}
}

func (r *NotLikeSpecification) IsEmpty() bool {
	return r.value == ""
}

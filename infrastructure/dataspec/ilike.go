package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type ILikeSpecification struct {
	field Field
	value string
}

func NewILike(field string, value string) dataset.Specifier {
	return &ILikeSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *ILikeSpecification) Joins(meta metadata.Meta) []string {
	return join(meta, r.field)
}

func (r *ILikeSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s ILIKE ?", r.field.ColumnName(meta))
}

func (r *ILikeSpecification) Values() []interface{} {
	return []interface{}{r.value}
}

func (r *ILikeSpecification) IsEmpty() bool {
	return r.value == ""
}

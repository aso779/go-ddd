package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type LikeSpecification struct {
	field Field
	value string
}

func NewLike(field string, value string) dataset.Specifier {
	return &LikeSpecification{
		field: NewField(field),
		value: value,
	}
}

func (r *LikeSpecification) Joins(meta metadata.Meta) []string {
	return join(meta, r.field)
}

func (r *LikeSpecification) Query(meta metadata.Meta) string {
	return fmt.Sprintf("%s LIKE ?", r.field.ColumnName(meta))
}

func (r *LikeSpecification) Values() []interface{} {
	return []interface{}{r.value}
}

func (r *LikeSpecification) IsEmpty() bool {
	return r.value == ""
}

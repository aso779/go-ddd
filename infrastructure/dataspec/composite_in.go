package dataspec

import (
	"fmt"
	"strings"

	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type CompositeInSpecification struct {
	fields []Field
	values any
}

func NewCompositeIn(fields []string, values any) dataset.Specifier {
	var fieldSet []Field
	for _, v := range fields {
		fieldSet = append(fieldSet, NewField(v))
	}

	return &CompositeInSpecification{
		fields: fieldSet,
		values: values,
	}
}

func (r *CompositeInSpecification) Joins(_ metadata.Meta) []metadata.Join {
	return []metadata.Join{}
}

func (r *CompositeInSpecification) Query(meta metadata.Meta) string {
	var columnNames []string
	for _, v := range r.fields {
		columnNames = append(columnNames, v.ColumnName(meta))
	}

	return fmt.Sprintf("(%s) IN (?)", strings.Join(columnNames, ","))
}

func (r *CompositeInSpecification) Values() []any {
	return []any{r.values}
}

func (r *CompositeInSpecification) IsEmpty() bool {
	return r.values == nil
}

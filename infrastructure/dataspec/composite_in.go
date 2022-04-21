package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"strings"
)

type CompositeInSpecification struct {
	fields []Field
	values []interface{}
}

func NewCompositeIn(fields []string, values []interface{}) dataset.Specifier {
	var fieldSet []Field
	for _, v := range fields {
		fieldSet = append(fieldSet, NewField(v))
	}

	return &CompositeInSpecification{
		fields: fieldSet,
		values: values,
	}
}

func (r *CompositeInSpecification) Joins(meta metadata.Meta) []string {
	return []string{}
}

func (r CompositeInSpecification) Query(meta metadata.Meta) string {
	var columnNames []string
	for _, v := range r.fields {
		columnNames = append(columnNames, v.ColumnName(meta))
	}

	return fmt.Sprintf("(%s) IN (?)", strings.Join(columnNames, ","))
}

func (r CompositeInSpecification) Values() []interface{} {
	return []interface{}{r.values}
}

func (r *CompositeInSpecification) IsEmpty() bool {
	return len(r.values) == 0
}

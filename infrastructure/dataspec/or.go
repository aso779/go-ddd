package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"strings"
)

type OrSpecification struct {
	specifications []dataset.Specifier
}

func NewOr(specifications ...dataset.Specifier) dataset.CompositeSpecifier {
	return &AndSpecification{
		specifications: specifications,
	}
}

func (r *OrSpecification) Joins(meta metadata.Meta) []string {
	uniqueIdx := make(map[string]struct{})
	var joins []string

	for _, specification := range r.specifications {
		for _, j := range specification.Joins(meta) {
			if _, ok := uniqueIdx[j]; !ok {
				joins = append(joins, j)
				uniqueIdx[j] = struct{}{}
			}
		}
	}

	return joins
}

func (r *OrSpecification) Query(meta metadata.Meta) string {
	var queries []string
	for _, specification := range r.specifications {
		queries = append(queries, specification.Query(meta))
	}

	query := strings.Join(queries, " OR ")

	return fmt.Sprintf("(%s)", query)
}

func (r *OrSpecification) Values() []interface{} {
	var values []interface{}
	for _, specification := range r.specifications {
		values = append(values, specification.Values()...)
	}
	return values
}

func (r *OrSpecification) Append(spec dataset.Specifier) {
	r.specifications = append(r.specifications, spec)
}

func (r *OrSpecification) IsNil() bool {
	return len(r.specifications) == 0
}
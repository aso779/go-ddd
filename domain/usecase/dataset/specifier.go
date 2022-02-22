package dataset

import "github.com/aso779/go-ddd/domain/usecase/metadata"

//Specifier specification interface
type Specifier interface {
	//Joins returns required joins slice
	Joins(meta metadata.Meta) []string
	//Query returns query condition string (where part)
	Query(meta metadata.Meta) string
	//Values returns query condition values
	Values() []interface{}
	//IsEmpty is empty flag
	IsEmpty() bool
}

type CompositeSpecifier interface {
	Specifier
	//Append additional specifications
	Append(specification Specifier)
}

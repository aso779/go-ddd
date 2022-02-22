package dataset

import "github.com/aso779/go-ddd/domain/usecase/metadata"

//Sorter sorter interface
type Sorter interface {
	//OrderBy return order by string
	OrderBy(meta metadata.Meta) string
	//IsEmpty is empty flag
	IsEmpty() bool
}

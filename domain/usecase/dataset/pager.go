package dataset

//Pager limit offset pager interface
type Pager interface {
	//GetSize page size
	GetSize() int
	//GetNumber page number
	GetNumber() int
	//GetOffset page offset
	GetOffset() int
	//IsEmpty is empty flag
	IsEmpty() bool
}

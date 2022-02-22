package dataset

//Pager limit offset pager interface
type Pager interface {
	//Size page size
	Size() int
	//Number page number
	Number() int
	//Offset page offset
	Offset() int
	//IsEmpty is empty flag
	IsEmpty() bool
}

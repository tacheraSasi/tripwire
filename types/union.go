package types

type UnionType interface {
	~string | ~int | ~float64 | ~float32 | ~bool
}
 
type Union[T UnionType] struct {
	Val *T
}

func New[T UnionType](){}
package common

type Tuple struct {
	First  interface{}
	Second interface{}
}

func NewTuple(first interface{}, second interface{}) *Tuple {
	return &Tuple{First: first, Second: second}
}

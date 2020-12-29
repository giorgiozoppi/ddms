package directory

import "github.com/golang-collections/collections/set"

// MultiValueDirectory maps a multiple value directory
type MultiValueDirectory interface {
	Get(key Key) ([]set.Set, error)
	Remove(key Key) ([]set.Set, error)
	GetSingle(key Key) (ValueInfo, error)
	PutSingle(key Key, value ValueInfo) error
	Next() (Key, ValueInfo, error)
	HasNext() bool
	Close() error
}

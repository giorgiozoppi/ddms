package directory

import (
	"github.com/giorgiozoppi/ddms/pkg/common"
	"github.com/golang-collections/collections/set"
)

// MultiValueDirectory maps a multiple value directory
type MultiValueDirectory interface {
	Get(key Key) ([]set.Set, error)
	Put(key Key, value ValueInfo) error
	Remove(key Key) ([]set.Set, error)
	GetSingle(key Key) (ValueInfo, error)
	Next() chan common.Tuple
	Close() error
}

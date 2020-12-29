package directory

// SingleValueDirectory maps a single value directory

type SingleValueDirectory interface {
	Get(key Key) (ValueInfo, error)
	Remove(key Key) (ValueInfo, error)
	Put(key Key, value ValueInfo) error
	Next() (Key, ValueInfo, error)
	HasNext() bool
	Close() error
}

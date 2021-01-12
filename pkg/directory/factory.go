package directory

import (
	"errors"
)

// Key is a directory key
type Key struct {
	Data []byte
}

// ValueInfo is a value to be stored in the directory
type ValueInfo struct {
	Data []byte
}

// Type is the kind of directory needed
type Type int32

// StorageType is a string that shows us a storage provider
type StorageType string

const (
	// RocksDbStorage is en enum value for indicating a rocksDB backend
	RocksDbStorage StorageType = rockDBProvider
)
const (
	// SingleDirectory is a single directory type
	SingleDirectory Type = iota
	// MultiDirectory is a multiple directory type
	MultiDirectory
)

// Configuration is the configuration for a directory service
type Configuration struct {
	Name        string
	FSystemPath string
	Type        Type
	BackendType StorageType
}

// NewProvider is a factory method for constructing a new provider.
func NewProvider(storageType StorageType, fsyncEnabled bool) (Provider, error) {
	switch storageType {
	case RocksDbStorage:
		return newRocksDBProvider(fsyncEnabled)
	}
	return nil, errors.New("invalid provider")
}

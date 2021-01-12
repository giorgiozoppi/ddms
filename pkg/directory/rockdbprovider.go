package directory

import (
	"sync"

	"github.com/golang-collections/collections/set"
	rocksdb "github.com/tecbot/gorocksdb"
)

// RocksDBProvider is a directory provider that uses RocksDB
type RocksDBProvider struct {
	mutex        sync.RWMutex
	options      *rocksdb.Options
	dataSource   *rocksdb.DB
	readOptions  *rocksdb.ReadOptions
	writeOptions *rocksdb.WriteOptions
	flushOptions *rocksdb.FlushOptions
}

const rockDBProvider = "RocksDB"

func newRocksDbMultipleValueDirectory(provider *RocksDBProvider, configuration Configuration) MultiValueDirectory {
	return &rocksDbMultiValueDirectory{dataSource: provider, configuration: configuration}
}

type rocksDbMultiValueDirectory struct {
	dataSource    *RocksDBProvider
	configuration Configuration
}

func (directory *rocksDbMultiValueDirectory) Put(key Key, value ValueInfo) error {
	return nil
}

func (directory *rocksDbMultiValueDirectory) Get(key Key) ([]set.Set, error) {
	data := make([]set.Set, 0)
	return data, nil
}
func (directory *rocksDbMultiValueDirectory) Remove(key Key) ([]set.Set, error) {
	var data []set.Set
	return data, nil
}
func (directory *rocksDbMultiValueDirectory) GetSingle(key Key) (ValueInfo, error) {
	return ValueInfo{}, nil
}
func (directory *rocksDbMultiValueDirectory) PutSingle(key Key, value ValueInfo) error {
	return nil
}
func (directory *rocksDbMultiValueDirectory) Next() chan struct {
	Key
	ValueInfo
} {
	//queue:=<-chan struct {Key; ValueInfo}{"http:...", 3}
	return nil
}

func (directory *rocksDbMultiValueDirectory) Close() error {
	return nil
}

// Create a single value directory provider
func newRocksDbSingleValueDirectory(handler *RocksDBProvider, configuration Configuration) SingleValueDirectory {
	return &rocksDbSingleValueDirectory{db: handler, configuration: configuration}
}

type rocksDbSingleValueDirectory struct {
	db            *RocksDBProvider
	configuration Configuration
}

func (directory *rocksDbSingleValueDirectory) Get(key Key) (ValueInfo, error) {

	return ValueInfo{}, nil
}
func (directory *rocksDbSingleValueDirectory) Remove(key Key) (ValueInfo, error) {
	return ValueInfo{}, nil
}
func (directory *rocksDbSingleValueDirectory) Put(key Key, value ValueInfo) error {
	var transactionDB *rocksdb.TransactionDBOptions
	transaction, status := rocksdb.OpenTransactionDb(directory.db.options, transactionDB, directory.configuration.FSystemPath)
	if status != nil {
		return nil
	}
	putError := transaction.Put(directory.db.writeOptions, key.Data, value.Data)
	transaction.Close()
	return putError
}
func (directory *rocksDbSingleValueDirectory) Next() (Key, ValueInfo, error) {
	return Key{}, ValueInfo{}, nil
}
func (directory *rocksDbSingleValueDirectory) HasNext() bool {
	return false
}
func (directory *rocksDbSingleValueDirectory) Close() error {
	return nil
}

// NewRocksDBProvider is a factory method for building a provider
func newRocksDBProvider(fsync bool) (Provider, error) {
	opts := rocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	ro := rocksdb.NewDefaultReadOptions()
	ro.SetFillCache(false)
	wo := rocksdb.NewDefaultWriteOptions()
	wo.SetSync(fsync)
	fo := rocksdb.NewDefaultFlushOptions()
	return &RocksDBProvider{
		options:      opts,
		readOptions:  ro,
		writeOptions: wo,
		flushOptions: fo,
	}, nil
}

// GetName is method that returns the name of the directory provider
func (db *RocksDBProvider) GetName() string {
	return rockDBProvider
}

// CreateSingleValueDirectory create a single value directory map
func (db *RocksDBProvider) CreateSingleValueDirectory(configuration Configuration) (SingleValueDirectory, error) {
	dataSource, err := rocksdb.OpenDb(db.options, configuration.FSystemPath)
	if err != nil {
		return nil, err
	}
	db.dataSource = dataSource
	singleValue := newRocksDbSingleValueDirectory(db, configuration)
	return singleValue, nil
}

// CreateMultipleValueDirectory create a multiple value directory map
func (db *RocksDBProvider) CreateMultipleValueDirectory(configuration Configuration) (MultiValueDirectory, error) {
	dataSource, err := rocksdb.OpenDb(db.options, configuration.FSystemPath)
	if err != nil {
		return nil, err
	}
	db.dataSource = dataSource
	multiValue := newRocksDbMultipleValueDirectory(db, configuration)
	return multiValue, nil
}

// RemoveDirectory removes a directory map
func (db *RocksDBProvider) RemoveDirectory(path string) error {
	return nil
}

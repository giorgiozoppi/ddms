package common

import (
	"errors"
	"sync"
)

// MaxBuckets is the maximum number of buckets in the hashmap
const MaxBuckets = uint64(128 * 1024)

// NodeMap is an hashtable of ID, Values
type NodeMap struct {
	entries  []*NodeList
	size     uint64
	mapMutex *sync.Mutex
}

// NewNodeMap creates a new node table
func NewNodeMap() *NodeMap {
	var mutexMap = new(NodeMap)
	mutexMap.entries = make([]*NodeList, MaxBuckets)
	mutexMap.mapMutex = new(sync.Mutex)
	return &NodeMap{
		entries:  make([]*NodeList, MaxBuckets),
		mapMutex: new(sync.Mutex),
	}
}

// Put puts a value into the hashtable
func (node *NodeMap) Put(key ID, value interface{}) error {
	item, errorItem := NewHashNode(key, value)
	if errorItem != nil {
		return errorItem
	}
	bucketIndex := item.Hash % MaxBuckets
	var errorInsert error
	if node.entries[bucketIndex] == nil {
		node.entries[bucketIndex] = NewNodeList(key, value)
	} else {
		_, errorInsert = node.entries[bucketIndex].Insert(key, value)
	}
	return errorInsert
}

// Get puts a value into the hashtable
func (node *NodeMap) Get(key ID) (*interface{}, error) {
	item, errorItem := NewHashKey(key)
	if errorItem != nil {
		return nil, errorItem
	}
	bucketIndex := item.Hash % MaxBuckets
	current, errorFind := node.entries[bucketIndex].Search(key)
	if current == nil {
		return nil, errorFind
	}
	return &current.Value, errorFind
}

// Remove removes the value from the hashtable
func (node *NodeMap) Remove(key ID) (interface{}, error) {
	node.mapMutex.Lock()
	defer node.mapMutex.Unlock()
	item, errorItem := NewHashKey(key)
	if errorItem != nil {
		return nil, errorItem
	}
	bucketIndex := item.Hash % MaxBuckets
	removeNode, errorFind := node.entries[bucketIndex].Remove(key)
	if removeNode == nil {
		return nil, errors.New("key not found")
	}
	return &removeNode.Value, errorFind
}

package common

import (
	"bytes"
	rand "crypto/rand"
	"errors"
	"github.com/minio/highwayhash"
	"sync"
)

// ID is a unique node identifier
type ID struct {
	Value []byte
}

// NewRandomID generate a random id for the table
func NewRandomID() (*ID, error) {
	data := make([]byte, 32)
	read, err := rand.Read(data)
	if err != nil {
		return nil, err
	}
	if read != 32 {
		return nil, errors.New("no so much entropy")
	}
	return &ID{
		Value: data,
	}, nil
}

// NodeList is a linked list concurrent
type NodeList struct {
	head      *HashNode
	listMutex sync.Mutex
	items     uint64
}

// HashNode is a node containing a key and value
type HashNode struct {
	Hash  uint64
	Key   ID
	Value interface{}
	Next  *HashNode
}

// NewHashNode is the factory method for a hash node
func NewHashNode(key ID, value interface{}) (*HashNode, error) {
	hash, err := highwayhash.New64(key.Value)
	if err != nil {
		return nil, err
	}
	hashValue := hash.Sum64()
	return &HashNode{
		Hash:  hashValue,
		Key:   key,
		Value: value,
		Next:  nil,
	}, nil
}
func NewHashKey(key ID) (*HashNode, error) {
	hash, err := highwayhash.New64(key.Value)
	if err != nil {
		return nil, err
	}
	hashValue := hash.Sum64()
	return &HashNode{
		Hash:  hashValue,
		Key:   key,
		Next:  nil,
	}, nil
}

// CompareTo compares an hash node to another node
func (node HashNode) CompareTo(o Comparable) int {
	tmp := o.(*HashNode)
	return bytes.Compare(node.Key.Value, tmp.Key.Value)
}

// NewNodeList creates a new node list
func NewNodeList(key ID, value interface{}) *NodeList {
	var list NodeList
	list.head = nil
	list.Insert(key, value)
	return &list
}

// Insert a key and value inside the list
func (list *NodeList) Insert(key ID, value interface{}) (*HashNode, error) {
	list.listMutex.Lock()
	defer list.listMutex.Unlock()
	var prev *HashNode
	var current *HashNode
	nodeCandidate, errCreation := NewHashNode(key, value)
	if errCreation!=nil {
		return nil, errCreation
	}
	prev = nil
	// empty list
	if list.head == nil {
		list.head = nodeCandidate
		return nodeCandidate, nil
	}
	for current=list.head; current.Next!=nil;  {
		if current.CompareTo(nodeCandidate) > 0 {
			// we have found the node.
			break
		}
		prev = current
		current = current.Next
	}
	if prev != nil {
		nodeCandidate.Next = current
		prev.Next = nodeCandidate
	} else {
		list.head = nodeCandidate
		nodeCandidate.Next = current
	}
	list.items =list.items + 1
	return nodeCandidate, nil
}

// Search a key in a node returns the previous node
func (list *NodeList) Search(key ID) (*HashNode,error) {
   node, _, errorSearch := list.searchKey(key)
   return node, errorSearch
}
func (list *NodeList) Clear() {
	for current:=list.head; current!=nil;  {
		current.Value = nil
		current = current.Next
	}
	list.head = nil
}
func (list NodeList) searchKey(key ID) (*HashNode, *HashNode,error){
	var prev *HashNode
	var current *HashNode
	prev = nil
	found := false
	for current=list.head; current!=nil && !found; {
		if bytes.Compare(current.Key.Value, key.Value) == 0 {
			found = true
			break
		}
		prev = current
		current = current.Next
	}
	if !found {
		return nil,nil, errors.New("item not found")
	}
	return current, prev, nil
}

// Remove an element from the list
func (list *NodeList) Remove(key ID) (*HashNode, error) {
	list.listMutex.Lock()
	defer list.listMutex.Unlock()
	node, prev, errorSearch := list.searchKey(key)
	if errorSearch != nil {
		return nil, errorSearch
	}
	if prev == nil {
		list.head = nil
		return node, nil
	} else {
		if node!=nil {
			prev.Next = node.Next
		} else {
			prev.Next = nil
		}
	}
	return node, errorSearch
}
func (list *NodeList) Iterator(abort <-chan struct{}) <-chan HashNode {
	ch := make(chan HashNode)
	go func() {
		defer close(ch)
		for ptr := list.head; ptr != nil; ptr=ptr.Next {
			select {
			case ch <- *ptr:
			case <-abort: // receive on closed channel can proceed immediately
				return
			}
		}
	}()
	return ch
}

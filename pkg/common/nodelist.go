package common

import (
	"bytes"
	"crypto/rand"
	"errors"
	"sync"

	"github.com/minio/highwayhash"
)

// ID is a unique node identifier
type ID struct {
	Value []byte
}

// ValueInfo is a list of values
type ValueInfo struct {
	Key   ID
	Value interface{}
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
}

// HashNode is a node containing a key and value
type HashNode struct {
	Hash  uint64
	Key   ID
	Value ValueInfo
	Next  *HashNode
}

// NewHashNode is the factory method for a hash node
func NewHashNode(key ID, value ValueInfo) (*HashNode, error) {
	hash, err := highwayhash.New64(key.Value)
	if err != nil {
		return nil, err
	}
	hashValue := hash.Sum64()
	value.Key = key
	return &HashNode{
		Hash:  hashValue,
		Key:   key,
		Value: value,
		Next:  nil,
	}, nil
}

// CompareTo compares an hash node to another node
func (node *HashNode) CompareTo(o Comparable) int {
	tmp := o.(*HashNode)
	return bytes.Compare(node.Key.Value, tmp.Key.Value)
}

// NewNodeList creates a new node list
func NewNodeList(key ID, value ValueInfo) *NodeList {
	var list NodeList
	head, _ := NewHashNode(key, value)
	list.head = head
	return &list
}

// Insert a key and value inside the list
func (list *NodeList) Insert(key ID, value ValueInfo) (*HashNode, error) {
	list.listMutex.Lock()
	defer list.listMutex.Unlock()
	nodeCandidate, errCreation := NewHashNode(key, value)
	current := list.head
	prev := list.head

	for {
		if current == nil {
			break
		}
		if current.CompareTo(nodeCandidate) > 0 {
			// we have found the node.
			break
		}
		prev = current
		current = current.Next
	}
	prev.Next = nodeCandidate
	nodeCandidate.Next = current
	return nodeCandidate, nil
}

// Search a key in a node returns the previuos node
func (list *NodeList) Search(key ID) (*HashNode, error) {
	current := list.head
	prev := list.head
	found := false
	for {
		if current == nil {
			break
		}
		if bytes.Compare(current.Key.Value, key.Value) == 0 {
			// we have found the node.
			found = true
			break
		}
		prev = current
		current = current.Next
	}
	if !found {
		return nil, errors.New("error not found")
	}
	return current, nil
}

// Remove an element from the list
func (list *NodeList) Remove(key ID) (*HashNode, error) {
	list.listMutex.Lock()
	defer list.listMutex.Unlock()
	current := list.head
	prev := list.head
	found := false
	for {
		if current == nil {
			break
		}
		if bytes.Compare(current.Key.Value, key.Value) == 0 {
			// we have found the node.
			found = true
			break
		}
		prev = current
		current = current.Next
	}
	if !found {
		return nil, errors.New("error not found")
	}
	prev.Next = current.Next
	return current, nil
}

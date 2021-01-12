package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeList(t *testing.T) {
	nodeID, _ := NewRandomID()
	nodeList := NewNodeList(*nodeID, 1)
	t.Run("should insert a node in a list", func(t *testing.T) {
		nodeID2, _ := NewRandomID()
		insertedNode, err := nodeList.Insert(*nodeID2, 2)
		require.NoError(t, err)
		require.NotNil(t, insertedNode)
		searchNode, errSearch := nodeList.Search(*nodeID2)
		require.NoError(t, errSearch)
		require.NotNil(t, searchNode)
		if searchNode != nil {
			require.Equal(t, searchNode.Key, *nodeID2)
		}
	})
	t.Run("should remove a id in a list", func(t *testing.T) {
		nodeID2, _ := NewRandomID()
		nodeList.Clear()
		insertedNode, err := nodeList.Insert(*nodeID2, 3)
		nodeList.Remove(*nodeID2)
		_, errSearch := nodeList.Search(*nodeID2)
		require.NoError(t, err)
		require.Error(t, errSearch)
		require.NotNil(t, insertedNode)

	})
	t.Run("should iterate for each item in the list", func(t *testing.T) {
		firstNode, _ := NewRandomID()
		secondNode, _ := NewRandomID()
		thirdNode, _ := NewRandomID()
		nodeList.Clear()
		nodeList.Insert(*firstNode, 2)
		nodeList.Insert(*secondNode, 10)
		nodeList.Insert(*thirdNode, 5)
		ch := make(chan struct{})
		for node := range nodeList.Iterator(ch) {
			require.True(t, (node.Value == 2) || (node.Value == 5) || (node.Value == 10))
		}
	})
}

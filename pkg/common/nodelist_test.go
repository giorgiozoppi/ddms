package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNodeList(t *testing.T) {
	nodeId, _ := NewRandomID()
	valueInfo:= 1
	nodeList:= NewNodeList(*nodeId, valueInfo)
	t.Run("should insert a node in a list", func(t *testing.T) {
		nodeId2, _ := NewRandomID()
		valueInfo2:= 2
		insertedNode, err := nodeList.Insert(*nodeId2, valueInfo2)
		require.NoError(t, err)
		require.NotNil(t, insertedNode)
		searchNode, errSearch := nodeList.Search(*nodeId2)
		require.NoError(t, errSearch)
		require.NotNil(t, searchNode)
		if searchNode!=nil {
			require.Equal(t, searchNode.Key, *nodeId2)
		}
	})

	/*
	t.Run("should remove a id in a list", func(t *testing.T) {
		nodeId2, _ := NewRandomID()
		valueInfo2:=ValueInfo{
			Value: valueInfo,
		}
		insertedNode, err := nodeList.Insert(*nodeId2, valueInfo2)
		nodeList.Remove(*nodeId2)
		nodeList.Search(*nodeId2)
		require.NoError(t, err)
		//require.NoError(t, errSearch)
		require.NotNil(t, insertedNode)
	})*/
}

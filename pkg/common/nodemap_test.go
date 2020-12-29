package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNodeMap(t *testing.T) {
	nodeMap:= NewNodeMap()
	randomId,_:=NewRandomID()

	t.Run("should insert a node in a map", func(t *testing.T) {
		nodeMap.Put(*randomId, 10)
		value, _:= nodeMap.Get(*randomId)
		intValue:=value.(int)
		require.Equal(t, 10,intValue)

	})
	t.Run("should remove a id in a map", func(t *testing.T) {
		_, errorType:=nodeMap.Remove(*randomId)
		require.NoError(t, errorType)
		_, errorValue:= nodeMap.Get(*randomId)
		require.Error(t, errorValue)
	})
}

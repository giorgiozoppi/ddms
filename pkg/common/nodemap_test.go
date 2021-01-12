package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeMap(t *testing.T) {
	nodeMap := NewNodeMap()
	randomID, _ := NewRandomID()

	t.Run("should insert a node in a map", func(t *testing.T) {
		nodeMap.Put(*randomID, 10)
		value, _ := nodeMap.Get(*randomID)
		intValue := value.(int)
		require.Equal(t, 10, intValue)

	})
	t.Run("should remove a id in a map", func(t *testing.T) {
		_, errorType := nodeMap.Remove(*randomID)
		require.NoError(t, errorType)
		_, errorValue := nodeMap.Get(*randomID)
		require.Error(t, errorValue)
	})
}

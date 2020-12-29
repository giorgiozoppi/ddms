package directory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRockDb(t *testing.T) {
	provider, error := NewProvider(RocksDbStorage, false)
	require.NoError(t, error, "provider not supported")
	t.Run("should fail when directory is invalid", func(t *testing.T) {
		_, singleValueError := provider.CreateSingleValueDirectory(Configuration{
			FSystemPath: "/23118298192-1212-1@",
		})
		require.Error(t, singleValueError, "creating a directory")
	})
}
func Test_ShouldInsertAndRetrieveCorrectlyInASingleValue(t *testing.T) {
	provider, error := NewProvider(RocksDbStorage, false)
	require.NoError(t, error, "provider not supported")
	directory, singleValueError := provider.CreateSingleValueDirectory(Configuration{
		FSystemPath: "/tmp",
	})
	require.NoError(t, singleValueError, "single value directory error")
	key := Key{
		Data: []byte("first"),
	}
	value := ValueInfo{
		Data: []byte("hello world"),
	}
	putError := directory.Put(key, value)
	require.NoError(t, putError, "put cannot fail")
	item, errorGet := directory.Get(key)
	require.NoError(t, errorGet, "get cannot fail")
	require.EqualValues(t, value, item, "key not inserted correctly")
}

func Test_ShouldInsertAndRetrieveCorrectlyInAMultipleValue(t *testing.T) {
	provider, error := NewProvider(RocksDbStorage, false)
	require.NoError(t, error, "provider not supported")
	directory, singleValueError := provider.CreateMultipleValueDirectory(Configuration{
		FSystemPath: "/tmp",
	})
	require.NoError(t, singleValueError, "single value directory error")
	key := Key{
		Data: []byte("first"),
	}
	value := ValueInfo{
		Data: []byte("hello world"),
	}

	putError := directory.Put(key, value)
	require.NoError(t, putError, "put cannot fail")
	item, errorGet := directory.Get(key)
	require.NoError(t, errorGet, "get cannot fail")
	require.EqualValues(t, value, item, "key not inserted correctly")
}

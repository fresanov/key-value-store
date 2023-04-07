package store

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	key   = "testKey"
	value = "testValue"
)

func TestPut(t *testing.T) {
	got, err := Get(key)
	require.Equal(t, "", got)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrorNoSuchKey)

	err = Put(key, value)
	require.NoError(t, err)

	got, err = Get(key)
	require.NoError(t, err)
	require.Equal(t, value, got)

	_ = Put(key, value)
	_ = Delete(key)
	got, err = Get(key)
	require.Equal(t, "", got)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrorNoSuchKey)
}

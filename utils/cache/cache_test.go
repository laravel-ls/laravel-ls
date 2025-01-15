package cache_test

import (
	"errors"
	"testing"

	"github.com/laravel-ls/laravel-ls/utils/cache"
	"github.com/stretchr/testify/require"
)

func TestCache_GetSet(t *testing.T) {
	c := cache.New[int]()
	c.Set("one", 1)
	c.Set("two", 2)

	v, ok := c.Get("one")
	require.Equal(t, 1, v)
	require.True(t, ok)

	v, ok = c.Get("three")
	require.Equal(t, 0, v)
	require.False(t, ok)
}

func TestCache_Remember(t *testing.T) {
	c := cache.New[int]()

	v, err := c.Remember("key", func(key string) (int, error) {
		require.Equal(t, key, "key")
		return 123, nil
	})

	require.Equal(t, v, 123)
	require.NoError(t, err)

	// Check that the value is stored in the cache
	v, ok := c.Get("key")

	require.Equal(t, v, 123)
	require.True(t, ok)

	// Test when error is returned
	v, err = c.Remember("key2", func(key string) (int, error) {
		require.Equal(t, key, "key2")
		return 332, errors.New("test error")
	})

	require.Equal(t, v, 332)
	require.Error(t, err)

	// Check that the value is not stored in cache.
	v, ok = c.Get("key2")

	require.Equal(t, v, 0)
	require.False(t, ok)
}

func TestCache_Forget(t *testing.T) {
	c := cache.New[int]()
	c.Set("first", 1234)
	c.Set("second", 999)

	v, ok := c.Get("first")
	require.Equal(t, 1234, v)
	require.True(t, ok)

	v, ok = c.Get("second")
	require.Equal(t, 999, v)
	require.True(t, ok)

	c.Forget("first")

	v, ok = c.Get("first")
	require.Equal(t, 0, v)
	require.False(t, ok)
}

func TestCache_Items(t *testing.T) {
	c := cache.New[string]()

	c.Set("one", "my_string")
	c.Set("two", "my_other_string")

	expected := map[string]string{
		"one": "my_string",
		"two": "my_other_string",
	}

	require.Equal(t, expected, c.Items())
}

func TestCache_Clear(t *testing.T) {
	c := cache.New[complex64]()

	c.Set("compl1", 0.123+1.234i)
	c.Set("compl2", 0.9271+32.8121i)
	c.Clear()

	require.Equal(t, 0, len(c.Items()))
}

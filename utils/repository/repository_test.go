package repository_test

import (
	"testing"

	"github.com/laravel-ls/laravel-ls/utils/repository"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Find(t *testing.T) {
	r := repository.Repository[int]{
		"one":  1234,
		"onex": 1234,
		"two":  999,
	}

	expected := map[string]int{
		"one":  1234,
		"onex": 1234,
	}

	assert.Equal(t, expected, r.Find("one"))
}

func TestRepository_Get(t *testing.T) {
	r := repository.Repository[string]{
		"key1": "value1",
		"key2": "value2",
	}

	v, ok := r.Get("key2")

	assert.Equal(t, "value2", v)
	assert.True(t, ok)

	v, ok = r.Get("key333")

	assert.Equal(t, "", v)
	assert.False(t, ok)
}

func TestRepository_Exists(t *testing.T) {
	r := repository.Repository[string]{
		"key1": "value1",
		"key2": "value2",
	}

	assert.True(t, r.Exists("key2"))
	assert.False(t, r.Exists("key323"))
}

func TestRepository_Clear(t *testing.T) {
	r := repository.Repository[float64]{
		"pi":     3.14,
		"random": 21.3,
	}

	r.Clear()

	assert.Equal(t, repository.Repository[float64]{}, r)
}

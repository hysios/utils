package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	m := NewMap[int, string]()

	assert.NotNil(t, m)

	m.Store(1, "abc")
	val, ok := m.Load(1)
	assert.True(t, ok)
	assert.Equal(t, val, "abc")
}

func TestLoadOrStore(t *testing.T) {
	m := NewMap[int, string]()

	assert.NotNil(t, m)

	val, ok := m.LoadOrStore(1, "abc")

	assert.False(t, ok)
	assert.Equal(t, val, "abc")
	m.Store(2, "abcd")
	val, ok = m.LoadOrStore(2, "abcd")
	assert.True(t, ok)
	assert.Equal(t, val, "abcd")
}

func TestLoadAndDelete(t *testing.T) {
	m := NewMap[int, string]()

	assert.NotNil(t, m)

	val, ok := m.LoadAndDelete(1)

	assert.False(t, ok)
	assert.Empty(t, val)
	m.Store(2, "abcd")
	val, ok = m.LoadAndDelete(2)
	assert.True(t, ok)
	assert.Equal(t, val, "abcd")

	val, ok = m.Load(2)
	assert.False(t, ok)
	assert.Empty(t, val)
}

func TestMapRange(t *testing.T) {
	m := NewMap[int, string]()

	assert.NotNil(t, m)

	m.Store(1, "abc")
	m.Store(2, "abcd")
	m.Store(3, "abcdef")

	var (
		keys   []int
		values []string
	)

	m.Range(func(key int, value string) bool {
		keys = append(keys, key)
		values = append(values, value)
		return true
	})

	assert.Contains(t, keys, 1)
	assert.Contains(t, keys, 2)
	assert.Contains(t, keys, 3)
	assert.Contains(t, values, "abc")
	assert.Contains(t, values, "abcd")
	assert.Contains(t, values, "abcdef")
}

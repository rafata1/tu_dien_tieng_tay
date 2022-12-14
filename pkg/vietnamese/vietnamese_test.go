package vietnamese

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	table := []struct {
		name     string
		array    []rune
		key      rune
		expected int
	}{
		{
			name:     "empty",
			array:    nil,
			key:      0,
			expected: -1,
		},
		{
			name:     "single-not-found",
			array:    []rune{'a'},
			key:      'b',
			expected: -1,
		},
		{
			name:     "single-found",
			array:    []rune{'b'},
			key:      'b',
			expected: 0,
		},
		{
			name:     "multiple-fond",
			array:    []rune{'a', 'b', 'c', 'd', 'e'},
			key:      'd',
			expected: 3,
		},
		{
			name:     "multiple-fond",
			array:    []rune{'a', 'b', 'c', 'd', 'e', 'e', 'f', 'f', 'g'},
			key:      'e',
			expected: 4,
		},
		{
			name:     "multiple-not-fond",
			array:    []rune{'a', 'b', 'c', 'd', 'f', 'f', 'g'},
			key:      'e',
			expected: -1,
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			result := binarySearch(e.array, e.key)
			assert.Equal(t, e.expected, result)
		})
	}
}

func TestRemoveAccent(t *testing.T) {
	result := RemoveAccent("Xin Chào Thế Giới")
	assert.Equal(t, "Xin Chao The Gioi", result)
}

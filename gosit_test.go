package gosit

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsValidNoteName(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected bool
	}{
		{"A", true},
		{"a", true},
		{"B", true},
		{"C", true},
		{"D", true},
		{"E", true},
		{"F", true},
		{"G", true},
		{"L", false},
		{"O", false},
		{"V", false},
		{"3", false},
	}

	for _, test := range tests {
		assert.Equal(IsValidNoteName(test.input), test.expected)
	}
}

package gosit

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestResolveNoteName(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected string
	}{
		{"a", "A"},
		{"A", "A"},
		{"a#", "A#/Bb"},
		{"ab", "G#/Ab"},
	}

	for _, test := range tests {
		res, err := ResolveNoteName(test.input)
		assert.Equal(err, nil)
		assert.Equal(res, test.expected)
	}
}

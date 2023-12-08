package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindValue(t *testing.T) {
	type test struct {
		input    string
		expected int
	}
	tests := []test{
		{
			input:    "two1nine",
			expected: 29,
		},
		{
			input:    "eightwothree",
			expected: 83,
		},
		{
			input:    "abcone2threexyz",
			expected: 13,
		},
		{
			input:    "xtwone3four",
			expected: 24,
		},
		{
			input:    "4nineeightseven2",
			expected: 42,
		},
		{
			input:    "zoneight234",
			expected: 14,
		},
		{
			input:    "7pqrstsixteen",
			expected: 76,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			require.Equal(t, tt.expected, findValue(tt.input))
		})
	}
}

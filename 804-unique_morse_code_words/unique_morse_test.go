package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMorse(t *testing.T) {
	var tests = []struct {
		id             int
		description    string
		input          []string
		expectedResult int
	}{
		{
			id:             1,
			description:    "From task description - 2 uniques",
			input:          []string{"gin", "zen", "gig", "msg"},
			expectedResult: 2,
		},
		{
			id:             2,
			description:    "one unique",
			input:          []string{"gin", "gin", "gin", "gin", "gin", "gin"},
			expectedResult: 1,
		},
		{
			id:             3,
			description:    "no translations",
			input:          nil,
			expectedResult: 0,
		},
		{
			id:             4,
			description:    "failed test from leetcode - two same words in array",
			input:          []string{"noilq", "kzlq", "ydreq", "ybxk", "kzlq"},
			expectedResult: 1,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", tc.id), func(t *testing.T) {
			actualResult := uniqueMorseRepresentations(tc.input)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}

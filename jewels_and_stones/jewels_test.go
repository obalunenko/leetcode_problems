package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJewels(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}

	var tests = []struct {
		id             int
		description    string
		input          args
		expectedResult int
	}{
		{
			id:          1,
			description: "One Jewel - lowercase",
			input: args{
				jewels: "a",
				stones: "AzalkK",
			},
			expectedResult: 1,
		},
		{
			id:          2,
			description: "3 Jewels - lowercase + uppercase",
			input: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			expectedResult: 3,
		},
		{
			id:          3,
			description: "Empty Jewels",
			input: args{
				jewels: "",
				stones: "aAAsdsdaK",
			},
			expectedResult: 0,
		},
		{
			id:          4,
			description: "Empty stones",
			input: args{
				jewels: "aA",
				stones: "",
			},
			expectedResult: 0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", tc.id), func(t *testing.T) {
			actualResult := numJewelsInStones(tc.input.jewels, tc.input.stones)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}

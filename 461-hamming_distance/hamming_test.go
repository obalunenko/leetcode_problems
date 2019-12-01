package hammingdistance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHamming(t *testing.T) {
	type args struct {
		a int
		b int
	}

	var tests = []struct {
		id             int
		description    string
		input          args
		expectedResult int
	}{
		{
			id:          1,
			description: "test case from task - different numbers",
			input: args{
				a: 1,
				b: 4,
			},
			expectedResult: 2,
		},
		{
			id:          2,
			description: "Diggerent length of binary",
			input: args{
				a: 1,
				b: 136,
			},
			expectedResult: 3,
		},
		{
			id:          3,
			description: "Same numbers",
			input: args{
				a: 136,
				b: 136,
			},
			expectedResult: 0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", tc.id), func(t *testing.T) {
			actualResult := hammingDistance(tc.input.a, tc.input.b)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}

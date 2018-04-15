package hamming_distance

import (
	"fmt"
	"testing"
)

type TestCase struct {
	id             int
	description    string
	input          Input
	expectedResult int
}

type Input struct {
	a int
	b int
}

type TestSuite []TestCase

var ts = TestSuite{
	TestCase{
		id:          1,
		description: "test case from task - different numbers",
		input: Input{
			a: 1,
			b: 4,
		},
		expectedResult: 2,
	},
	TestCase{
		id:          2,
		description: "Diggerent length of binary",
		input: Input{
			a: 1,
			b: 136,
		},
		expectedResult: 3,
	},
	TestCase{
		id:          3,
		description: "Same numbers",
		input: Input{
			a: 136,
			b: 136,
		},
		expectedResult: 0,
	},
}

func TestHamming(t *testing.T) {

	for _, tc := range ts {
		t.Run(fmt.Sprintf("Test%d", tc.id), func(t *testing.T) {

			actualResult := hammingDistance(tc.input.a, tc.input.b)

			if actualResult != tc.expectedResult {
				fmt.Printf(" - Got: %d\n - Expected: %d\n", actualResult, tc.expectedResult)
				t.Fail()
			}

		})
	}

}

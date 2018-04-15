package judge_route_circle

import (
	"fmt"
	"testing"
)

type TestCase struct {
	id             int
	description    string
	input          Input
	expectedResult ExpectedResult
}

type Input struct {
	moves string
}

type ExpectedResult struct {
	circle bool
}

type TestSuite []TestCase

var ts = TestSuite{
	TestCase{
		id:          1,
		description: "test case from task - returned to start",
		input: Input{
			moves: "UD",
		},
		expectedResult: ExpectedResult{
			circle: true,
		},
	},
	TestCase{
		id:          2,
		description: "test case from task - not returned to start",
		input: Input{
			moves: "LL",
		},
		expectedResult: ExpectedResult{
			circle: false,
		},
	},
	TestCase{
		id:          3,
		description: "Not returned to start",
		input: Input{
			moves: "ULUUDRLD",
		},
		expectedResult: ExpectedResult{
			circle: false,
		},
	},
	TestCase{
		id:          4,
		description: "Not returned to start on y",
		input: Input{
			moves: "ULUDUDRLD",
		},
		expectedResult: ExpectedResult{
			circle: false,
		},
	},
	TestCase{
		id:          5,
		description: "Not returned to start on x",
		input: Input{
			moves: "ULUUDRLDR",
		},
		expectedResult: ExpectedResult{
			circle: false,
		},
	},
	TestCase{
		id:          6,
		description: "Returned to start after Long walk",
		input: Input{
			moves: "ULUUDRLDRD",
		},
		expectedResult: ExpectedResult{
			circle: true,
		},
	},
}

func TestHamming(t *testing.T) {

	for _, tc := range ts {
		t.Run(fmt.Sprintf("Test%d", tc.id), func(t *testing.T) {

			actualResult := judgeCircle(tc.input.moves)

			if actualResult != tc.expectedResult.circle {
				fmt.Printf(" - Got: %v\n - Expected: %v\n", actualResult, tc.expectedResult.circle)
				t.Fail()
			}

		})
	}

}

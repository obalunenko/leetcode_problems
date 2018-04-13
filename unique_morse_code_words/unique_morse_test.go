package unique_morse_code_words

import (
	"fmt"
	"testing"
)

type TestCase struct {
	id             int
	description    string
	input          []string
	expectedResult int
}

type TestSuite []TestCase

var ts = TestSuite{
	TestCase{
		id:             1,
		description:    "From task description - 2 uniqs",
		input:          []string{"gin", "zen", "gig", "msg"},
		expectedResult: 2,
	},
}

func TestMorse(t *testing.T) {

	for _, tc := range ts {
		t.Run(fmt.Sprintf("Test%d", tc.id), func(t *testing.T) {
			actualResult := uniqueMorseRepresentations(tc.input)
			if actualResult != tc.expectedResult {
				fmt.Printf(" - Got: %d\n - Expected: %d\n", actualResult, tc.expectedResult)
				t.Fail()
			}
		})
	}
}

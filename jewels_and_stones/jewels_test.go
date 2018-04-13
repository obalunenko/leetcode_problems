package jewels_and_stones

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
	jewels string
	stones string
}

type TestSuite []TestCase

var ts = TestSuite{
	TestCase{
		id:          1,
		description: "One Jewel - lowercase",
		input: Input{
			jewels: "a",
			stones: "AzalkK",
		},
		expectedResult: 1,
	},
	TestCase{
		id:          2,
		description: "3 Jewels - lowercase + uppercase",
		input: Input{
			jewels: "aA",
			stones: "aAAbbbb",
		},
		expectedResult: 3,
	},
	TestCase{
		id:          3,
		description: "Empty Jewels",
		input: Input{
			jewels: "",
			stones: "aAAsdsdaK",
		},
		expectedResult: 0,
	},
	TestCase{
		id:          4,
		description: "Empty stones",
		input: Input{
			jewels: "aA",
			stones: "",
		},
		expectedResult: 0,
	},
}

func TestJewels(t *testing.T) {

	for _, tc := range ts {
		t.Run(fmt.Sprintf("Test%d", tc.id), func(t *testing.T) {

			actualResult := numJewelsInStones(tc.input.jewels, tc.input.stones)

			if actualResult != tc.expectedResult {
				fmt.Printf(" - Got: %d\n - Expected: %d\n", actualResult, tc.expectedResult)
				t.Fail()
			}

		})
	}

}

package jewels_and_stones

import "fmt"

func numJewelsInStones(J string, S string) int {
	var jewelsCount int

	for _, itemS := range S {
		oneS := string(itemS)
		for _, itemJ := range J {
			oneJ := string(itemJ)
			if oneJ == oneS {
				fmt.Printf("Jewel found! [%s] equals [%s]", oneJ, oneS)
				jewelsCount++
			}
		}

	}

	return jewelsCount
}

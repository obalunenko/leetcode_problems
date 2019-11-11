package main

import "fmt"

func numJewelsInStones(j string, s string) int {
	var jewelsCount int

	for _, itemS := range s {
		oneS := string(itemS)

		for _, itemJ := range j {
			oneJ := string(itemJ)

			if oneJ == oneS {
				fmt.Printf("Jewel found! [%s] equals [%s]", oneJ, oneS)
				jewelsCount++
			}
		}
	}

	return jewelsCount
}

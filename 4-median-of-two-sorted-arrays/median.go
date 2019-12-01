package median

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, len(nums1)+len(nums2))
	merged = append(nums1, nums2...)

	sort.Ints(merged)

	fmt.Printf("merged: %+v \n", merged)

	var median float64
	isbalanced := len(merged) % 2
	if isbalanced == 0 {
		middle := len(merged) / 2.0
		fmt.Printf("middle [%v]\n", middle)

		median = float64(merged[middle]+merged[middle-1]) / 2
	} else {
		middle := (len(merged) - 1) / 2
		fmt.Printf("middle [%v]\n", middle)
		median = float64(merged[(middle)])
	}
	fmt.Printf("median: %v \n", median)
	return median
}

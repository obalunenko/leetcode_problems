package hamming_distance

import (
	"fmt"
	"strconv"
)

func hammingDistance(x int, y int) int {
	var result int
	binX := strconv.FormatInt(int64(x), 2)
	binY := strconv.FormatInt(int64(y), 2)

	fmt.Printf("Before diff fix:\n - X: %s\n - Y: %s\n", binX, binY)
	var diffLen int

	switch {
	case len(binX) < len(binY):
		diffLen = len(binY) - len(binX)
		binX = addLeadZeros(binX, diffLen)

	case len(binX) > len(binY):
		diffLen = len(binX) - len(binY)
		binY = addLeadZeros(binY, diffLen)
	}

	fmt.Printf("After diff fix\n - X: %s\n - Y: %s\n", binX, binY)
	for index, digit := range binX {
		if string([]rune(binY)[index]) != string(digit) {

			result++

		}
	}

	return result
}

func addLeadZeros(bin string, diff int) string {
	//fmt.Printf("Diff is %d\n", diff)
	//fmt.Printf("string is %s\n", bin)

	var tmp string
	for i := 0; i < diff; i++ {
		tmp += "0"
		//fmt.Printf("Step[%d]: tmp is [%s]\n", i, tmp)
	}
	bin = tmp + bin
	return bin
}

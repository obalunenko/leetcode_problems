package main

import (
	"strconv"
)

func hammingDistance(x int, y int) int {
	var (
		result  int
		diffLen int
	)

	binX := strconv.FormatInt(int64(x), 2)
	binY := strconv.FormatInt(int64(y), 2)

	if len(binX) < len(binY) {
		diffLen = len(binY) - len(binX)
		binX = addLeadZeros(binX, diffLen)
	} else {
		diffLen = len(binX) - len(binY)
		binY = addLeadZeros(binY, diffLen)
	}

	for index, digit := range binX {
		if string([]rune(binY)[index]) != string(digit) {
			result++
		}
	}

	return result
}

func addLeadZeros(bin string, diff int) string {
	var tmp string
	for i := 0; i < diff; i++ {
		tmp += "0"
	}

	return tmp + bin
}

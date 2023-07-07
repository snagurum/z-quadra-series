package main

import (
	"fmt"
	"math"
)

func main() {
	n := 32
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%4d", some(i, j))
		}
		fmt.Println()
	}
	// fmt.Println(some(1, 1), " ", some(1, 2), " ", some(2, 1), " ", some(2, 2))
	// fmt.Println(some(1, 3), some(1, 4), some(3, 1), some(2, 3))
	// fmt.Println(some(1, 5))
}

func some(r, c int) int {
	if r == 1 && c == 1 {
		return 1
	}
	if r == 1 && c == 2 {
		return 2
	}
	if r == 2 && c == 1 {
		return 3
	}
	if r == 2 && c == 2 {
		return 4
	}

	newR := 0
	newC := 0
	exponent := getQuadra(int(math.Max(float64(r), float64(c))))
	_4PowerVal := int(math.Pow(float64(4), float64(exponent))) //4, 16, 64, 128
	_2PowerVal := int(math.Pow(float64(2), float64(exponent))) //0,  2,  4,   8, 16

	if _2PowerVal < r && _2PowerVal < c {
		newR = r - _2PowerVal
		newC = c - _2PowerVal
		_4PowerVal = 3 * _4PowerVal
	} else if _2PowerVal < r {
		newR = r - _2PowerVal
		newC = c
		_4PowerVal = 2 * _4PowerVal
	} else if _2PowerVal < c {
		newC = c - _2PowerVal
		newR = r
	} else {
		return 0
	}

	return _4PowerVal + some(newR, newC)
}

func getQuadra(c int) int {
	num := 0
	for int(math.Pow(float64(4), float64(num))) <= (c-1)*(c-1) {
		num++
	}
	return num - 1
}

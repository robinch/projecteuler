package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findLargestNdigitPal(3))
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	return s == strRevers(s)
}

func strRevers(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func findLargestNdigitPal(n int) int {
	maxRes := 0

	step := 1

	min := 1
	for i := 0; i < n-1; i++ {
		min *= 10
	}
	max := min*10 - 1

	for i := max; i >= min; i = i - step {
		for j := max; j >= min; j = j - step {
			if isPalindrome(i * j) {
				if i*j > maxRes {
					maxRes = i * j
				}
			}
		}
	}
	return maxRes
}

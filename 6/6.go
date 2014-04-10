package main

// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385

// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 552 = 3025

// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()
	n, err := strconv.Atoi(flag.Arg(0))
	if err == nil {
		fmt.Println(squareOfSum(n) - sumOfSquare(n))
	} else {
		fmt.Println("You need an argument")
	}

}

func squareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

func sumOfSquare(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

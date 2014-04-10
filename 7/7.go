package main

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

func main() {
	flag.Parse()
	n, err := strconv.Atoi(flag.Arg(0))
	if err == nil {
		fmt.Println(getNumberOfPrimes(n))
	} else {
		fmt.Println("You need one int argument")
	}
}

func getNumberOfPrimes(n int) int {
	nrOfPrimes := 0
	maxPrime := 0
	for i := 2; nrOfPrimes < n; i++ {
		if isPrime(i) {
			nrOfPrimes++
			maxPrime = i
		}
	}
	return maxPrime
}

func isPrime(n int) bool {
	sqrt := int(math.Ceil(math.Sqrt(float64(n))))
	if n == 2 {
		return true
	}
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

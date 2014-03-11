package main

import (
	"fmt"
	"math"
)

// Smallest multiple
// Problem 5
// 2520 is the smallest number that can be divided
// by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is
// evenly divisible by all of the numbers from 1 to 20?

func main() {

	fmt.Println("Test:", calcIntFromPrimeMap(getLcmPrimeMap(10)))
	fmt.Println("Real deal:", calcIntFromPrimeMap(getLcmPrimeMap(20)))
}

func calcIntFromPrimeMap(m map[int]int) int {
	res := 1
	for key, value := range m {
		for i := 0; i < value; i++ {
			res *= key
		}
	}

	return res
}

func getLcmPrimeMap(n int) map[int]int {
	var lcmPrimeMap map[int]int
	lcmPrimeMap = make(map[int]int)

	for i := 1; i < n+1; i++ {
		primes := getPrimeMap(i)
		for key, value := range primes {
			if lcmPrimeMap[key] < value {
				lcmPrimeMap[key] = value
			}
		}
	}

	return lcmPrimeMap
}

func getPrimeMap(n int) map[int]int {
	var m map[int]int
	m = make(map[int]int)
	sqrt := int(math.Ceil(math.Sqrt(float64(n))))

	for i := 2; i < sqrt+1; i++ {
		for n%i == 0 {
			m[i]++
			n = n / i
		}
	}
	if n != 1 {
		m[n]++
	}
	return m
}

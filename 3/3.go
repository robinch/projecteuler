package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(naive(600851475143))
}

func naive(n int) []int {
	var list []int
	sqrt := int(math.Ceil(math.Sqrt(float64(n))))

	for i := 2; i < sqrt; i++ {
		for n%i == 0 {
			list = append(list, i)
			n = n / i
		}
	}
	if n != 1 {
		list = append(list, n)
	}
	return list
}

// func fermat(n int){
// 	var list []int
// 	for n%2 == 0 {
// 		append(list, 2)
// 	}
// }

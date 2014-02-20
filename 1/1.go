package main

import "fmt"

func main() {
	fmt.Printf("%d", multiples_3_5(1000))
}

func multiples_3_5(n int) int {
	total := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
	}
	return total
}

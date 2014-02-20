package main

import "fmt"

func main() {
	MAX := 4000000
	sumEven := 0
	fibReturn := 0
	fib := generateFib()

	for {
		fibReturn = fib()
		if fibReturn > MAX {
			fmt.Println(sumEven)
			return
		} else if fibReturn%2 == 0 {
			sumEven += fibReturn
		}
	}

}

func generateFib() func() int {
	a := 0
	b := 1
	return func() int {
		b = a + b
		a = b - a
		return b
	}
}

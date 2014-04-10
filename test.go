package main

import "fmt"

func main() {
	price := 4000.0
	for i := 0; i < 30; i++ {
		price = price * 1.02
	}

	fmt.Println(price)
}
